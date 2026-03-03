package event

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	nc *nats.Conn
}

func NewSubscriber(nc *nats.Conn) *Subscriber {
	return &Subscriber{nc: nc}
}

func (s *Subscriber) Subscribe(subject string, handler nats.MsgHandler) error {
	_, err := s.nc.Subscribe(subject, handler)
	if err != nil {
		log.Printf("Error subscribing to subject %s: %v", subject, err)
		return err
	}
	return nil
}

func (s *Subscriber) SubscribeContext(ctx context.Context, subject string, handler nats.MsgHandler) error {
	sub, err := s.nc.Subscribe(subject, handler)
	if err != nil {
		log.Printf("Error subscribing to subject %s: %v", subject, err)
		return err
	}

	go func() {
		<-ctx.Done()
		if err := sub.Unsubscribe(); err != nil {
			log.Printf("Error unsubscribing from subject %s: %v", subject, err)
		}
	}()

	return nil
}