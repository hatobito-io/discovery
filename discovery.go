package discovery

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	dproto "github.com/hatobito-io/discovery/proto"
	"github.com/nats-io/nats.go"
)

// DefaultSubjectPrefix is the default subject prefix for messages sent by
// service discovery agent.
const DefaultSubjectPrefix = "github.com.hatobitoio.discovery"

// DefaultUpdateInterval is the default update interval for agents. If an agent
// does not send an update within expected update interval, other agents will
// forget about services registered by offending agent. 10% threshold is always
// added.
const DefaultUpdateInterval = time.Millisecond * 1000

// Agent is a service discovery agent.
type Agent struct {
	conn             *nats.Conn
	subjectPrefix    string
	prefixParts      int
	knownServices    *infoList
	providedServices *infoList
	subs             *nats.Subscription
	watched          map[string]bool
	connected        bool
	running          bool
	l                chan struct{}
	send             chan *msgWrapper
	clientID         string
}

func (a *Agent) lock() {
	a.l <- struct{}{}
}

func (a *Agent) unlock() {
	<-a.l
}

// NewAgent creates a new service discovery agent.
func NewAgent(conn *nats.Conn, opts ...Option) (*Agent, error) {
	s := &Agent{
		conn:             conn,
		subjectPrefix:    DefaultSubjectPrefix,
		knownServices:    &infoList{},
		providedServices: &infoList{},
		watched:          make(map[string]bool),
		l:                make(chan struct{}, 1),
	}
	var cid [16]byte
	if _, err := rand.Read(cid[:]); err != nil {
		return nil, err
	}

	s.prefixParts = len(strings.Split(s.subjectPrefix, "."))
	s.clientID = hex.EncodeToString(cid[:])
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	go s.ConnStateHandler(conn)
	return s, nil
}

// Register registers a service instance making it available
// for discovery by other services.
func (a *Agent) Register(info *ServiceInfo) error {
	a.lock()
	defer a.unlock()
	item := a.providedServices.find(info)
	if item != nil {
		item.Address = info.Address
		item.Name = info.Name
	} else {
		a.providedServices.insert(info)
	}
	if a.running && a.connected {
		a.send <- &msgWrapper{
			subject: a.serviceListSubject(),
			msg: &dproto.ServicesList{Services: []*dproto.ServiceInfoProto{
				{
					Address: info.Address,
					Name:    info.Name,
				},
			}},
		}
	}

	return nil
}

// Unregister removes a service instance making it unavailable
// for discovery by other services.
func (a *Agent) Unregister(info *ServiceInfo) error {
	if item := a.providedServices.find(info); item != nil {
		a.providedServices.remove(item)
	}
	return nil
}

// Start starts the discovery service
func (a *Agent) Start() error {
	a.lock()
	defer a.unlock()
	if a.running {
		return errors.New("discovery agent is already running")
	}
	subs, err := a.conn.Subscribe(a.subjectPrefix+".>", a.handleMessage)
	if err != nil {
		return err
	}
	a.subs = subs
	a.running = true
	if a.connected {
		a.startStopWorker(true)
	}
	return nil
}

// Stop stops the discovery service. It will notify other instances that all
// services registered by this agent instance are no longer available. The agent
// will forget all known remote services immediately. Local services list is
// kept intact.
func (a *Agent) Stop() error {
	a.lock()
	defer a.unlock()
	if !a.running {
		return errors.New("discovery agent is not running")
	}
	a.running = false
	a.subs.Unsubscribe()
	a.subs = nil
	a.knownServices.clear()
	if a.connected {
		a.publish(a.stopSubject(), nil)
		a.startStopWorker(false)
	}
	return nil
}

// Watch expresses interest in particular service. Only watched services will be
// available for discovery.
func (a *Agent) Watch(serviceName string) error {
	a.lock()
	defer a.unlock()
	if a.watched[serviceName] {
		return nil
	}
	a.watched[serviceName] = true
	if a.connected && a.running {
		msg := &dproto.ServiceInterest{ServiceName: []string{serviceName}}
		a.publish(a.interestSubject(), msg)
	}
	return nil
}

// Unwatch stops monitoring of particular service availability.
func (a *Agent) Unwatch(serviceName string) {
	a.lock()
	defer a.unlock()
	item := a.knownServices.first
	for item != nil {
		next := item.next
		if item.Name == serviceName {
			a.knownServices.remove(item)
		}
		item = next
	}
	delete(a.watched, serviceName)
}

// Discover returns a list of last known addresses of a service. If includeLocal
// is false, the services registered by this instance of Agent using Register()
// will be omitted.
func (a *Agent) Discover(serviceName string, includeLocal bool) []*ServiceInfo {
	var ret []*ServiceInfo
	var lists []*infoList
	if includeLocal {
		lists = []*infoList{a.knownServices, a.providedServices}
	} else {
		lists = []*infoList{a.knownServices}
	}
	for _, list := range lists {
		s := list.first
		for s != nil {
			if s.Name == serviceName {
				item := *s
				ret = append(ret, &item)
			}
			s = s.next
		}
	}
	return ret
}

// ConnStateHandler is used to monitor the state of NATS connection. This method
// should be called when NATS connection state changes: connect, disconnect,
// close. Not calling this method after connection state change will result in
// other agents losing the knowledge about services registered by this instance
// of Agent.
func (a *Agent) ConnStateHandler(conn *nats.Conn) {
	a.lock()
	defer a.unlock()
	connected := conn.IsConnected()
	if connected == a.connected {
		return
	}
	a.connected = connected
	if a.running {
		a.startStopWorker(connected)
	}
}

func (a *Agent) startStopWorker(start bool) {
	if start {
		a.send = make(chan *msgWrapper, 10)
		go worker(a)
		watchedServices := make([]string, len(a.watched))
		for serviceName := range a.watched {
			watchedServices = append(watchedServices, serviceName)
		}
		a.send <- &msgWrapper{
			subject: a.interestSubject(),
			msg:     &dproto.ServiceInterest{ServiceName: watchedServices},
		}
	} else {
		close(a.send)
		a.send = nil
	}
}
