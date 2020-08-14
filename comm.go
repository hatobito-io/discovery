package discovery

import (
	"time"

	"github.com/golang/protobuf/proto"
	dproto "github.com/hatobito-io/discovery/proto"
)

type msgWrapper struct {
	subject string
	msg     proto.Message
}

func (a *Agent) publish(subject string, msg proto.Message) {
	a.send <- &msgWrapper{subject: subject, msg: msg}
}

func worker(a *Agent) {
	ch := a.send
	timer := time.NewTicker(DefaultUpdateInterval)
	a.sendUpdates()
	a.checkExpiration()
receiving:
	for {
		select {
		case msg := <-ch:
			if msg == nil {
				break receiving
			}
			a.publishMessage(msg)
		case <-timer.C:
			a.sendUpdates()
			a.checkExpiration()
		}
	}
	timer.Stop()
}

func (a *Agent) publishMessage(msg *msgWrapper) error {
	var data []byte
	if msg.msg != nil {
		var err error
		data, err = proto.Marshal(msg.msg)
		if err != nil {
			return err
		}
	}
	a.conn.Publish(msg.subject, data)
	return nil
}

func (a *Agent) checkExpiration() error {
	a.lock()
	defer a.unlock()
	now := time.Now()
	item := a.knownServices.first
	for item != nil {
		next := item.next
		if now.After(item.GoodUntil) {
			a.knownServices.remove(item)
		}
		item = next
	}
	return nil
}

func (a *Agent) sendUpdates() error {
	a.lock()
	defer a.unlock()
	item := a.providedServices.first
	msg := &dproto.ServicesList{}
	for item != nil {
		s := &dproto.ServiceInfoProto{
			Address:  item.Address,
			ClientId: a.clientID,
			Name:     item.Name,
		}
		msg.Services = append(msg.Services, s)
		item = item.next
	}
	if len(msg.Services) > 0 {
		a.publishMessage(&msgWrapper{
			subject: a.serviceListSubject(),
			msg:     msg,
		})
	}
	return nil
}
