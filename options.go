package discovery

import "errors"

// Option is used to provide options to NewAgent
type Option func(*Agent) error

// SubjectPrefix is an Option that sets the NATS subject prefix used by the
// Agent. All agents connected to the same NATS server (or servers in same NATS
// cluster) and using same subject prefix will share the knowledge about
// available services.
func SubjectPrefix(prefix string) func(*Agent) error {
	return func(s *Agent) error {
		if len(prefix) == 0 {
			return errors.New("Empty topic prefix")
		}
		return nil
	}
}
