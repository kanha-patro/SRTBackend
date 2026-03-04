package event

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type EventPublisher struct {
	nc *nats.Conn
}

func NewEventPublisher(nc *nats.Conn) *EventPublisher {
	return &EventPublisher{nc: nc}
}

func (ep *EventPublisher) Publish(subject string, data interface{}) error {
	message, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling event data: %v", err)
		return err
	}

	if err := ep.nc.Publish(subject, message); err != nil {
		log.Printf("Error publishing event to NATS: %v", err)
		return err
	}

	return nil
}

// Publisher is the minimal interface for publishing events.
type Publisher interface {
	Publish(subject string, data interface{}) error
}
