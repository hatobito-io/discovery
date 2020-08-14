package discovery

import (
	"strings"

	dproto "github.com/hatobito-io/discovery/proto"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func (a *Agent) serviceListSubject() string {
	return a.subjectPrefix + ".servicelist." + a.clientID
}

func (a *Agent) stopSubject() string {
	return a.subjectPrefix + ".stop." + a.clientID
}

func (a *Agent) interestSubject() string {
	return a.subjectPrefix + ".interest." + a.clientID
}

func (a *Agent) parseNatsMessage(msg *nats.Msg) (proto.Message, string, bool) {
	parts := strings.Split(msg.Subject, ".")
	nParts := len(parts)
	matched := nParts > a.prefixParts+1 &&
		strings.Join(parts[:a.prefixParts], ".") == a.subjectPrefix
	if !matched {
		return nil, "", false
	}
	action := parts[a.prefixParts]
	clientID := strings.Join(parts[a.prefixParts+1:], ".")
	myself := clientID == a.clientID
	var result proto.Message
	switch action {
	case "interest":
		result = &dproto.ServiceInterest{}
	case "servicelist":
		result = &dproto.ServicesList{}
	case "stop":
		result = &dproto.AgentStopped{}
	}
	if result == nil {
		return nil, clientID, myself
	}
	if err := proto.Unmarshal(msg.Data, result); err != nil {
		return nil, clientID, myself
	}
	return result, clientID, myself
}
