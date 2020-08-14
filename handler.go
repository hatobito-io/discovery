package discovery

import (
	"time"

	dproto "github.com/hatobito-io/discovery/proto"
	"github.com/nats-io/nats.go"
)

func (a *Agent) handleMessage(msg *nats.Msg) {
	decoded, clientID, myself := a.parseNatsMessage(msg)
	if decoded == nil {
		return
	}
	switch decoded := decoded.(type) {
	case *dproto.ServiceInterest:
		if !myself {
			a.handleInterestMessage(decoded, clientID)
		}
	case *dproto.ServicesList:
		if !myself {
			a.handleServiceListMessage(decoded, clientID)
		}
	case *dproto.AgentStopped:
		if !myself {
			a.handleStopMessage(clientID)
		}
	}
}

func (a *Agent) handleStopMessage(clientID string) {
	a.lock()
	defer a.unlock()
	item := a.knownServices.first
	for item != nil {
		next := item.next
		if item.updatedBy == clientID {
			a.knownServices.remove(item)
		}
		item = next
	}
}

func (a *Agent) handleServiceListMessage(msg *dproto.ServicesList, clientID string) {
	now := time.Now()
	deadline := now.Add(DefaultUpdateInterval + DefaultUpdateInterval/10)
	if len(msg.Services) < 1 {
		return
	}
	a.lock()
	defer a.unlock()
	for _, svc := range msg.Services {
		search := &ServiceInfo{
			Address:   svc.Address,
			Name:      svc.Name,
			updatedBy: clientID,
			UpdatedAt: now,
			GoodUntil: deadline,
		}
		if item := a.knownServices.find(search); item != nil {
			item.updatedBy = clientID
			item.UpdatedAt = now
			item.GoodUntil = deadline
		} else {
			a.knownServices.insert(search)
		}
	}

}

func (a *Agent) handleInterestMessage(msg *dproto.ServiceInterest, clientID string) {
	if len(msg.ServiceName) < 1 {
		return
	}
	a.lock()
	defer a.unlock()
	item := a.providedServices.first
	reply := &dproto.ServicesList{}
	for item != nil {
		if contains(msg.ServiceName, item.Name) {
			s := &dproto.ServiceInfoProto{
				Address:  item.Address,
				ClientId: a.clientID,
				Name:     item.Name,
			}
			reply.Services = append(reply.Services, s)
		}
		item = item.next
	}
	if len(reply.Services) > 0 {
		a.send <- &msgWrapper{
			subject: a.serviceListSubject(),
			msg:     reply,
		}
	}
}
