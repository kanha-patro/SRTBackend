package nats

import (
	"time"

	"github.com/nats-io/nats.go"
)

type NATSClient struct {
	conn *nats.Conn
}

// NewNATSClient initializes a new NATS client
func NewNATSClient(url string) (*NATSClient, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NATSClient{conn: nc}, nil
}

// Publish sends a message to a specified subject
func (c *NATSClient) Publish(subject string, msg []byte) error {
	return c.conn.Publish(subject, msg)
}

// Subscribe subscribes to a specified subject and processes messages with the provided handler
func (c *NATSClient) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.Subscribe(subject, handler)
}

// Close closes the NATS connection
func (c *NATSClient) Close() {
	c.conn.Close()
}

// WaitForReady waits until the NATS connection is ready
func (c *NATSClient) WaitForReady(timeout time.Duration) error {
	start := time.Now()
	for {
		if c.conn.IsConnected() {
			return nil
		}
		if time.Since(start) > timeout {
			return nats.ErrTimeout
		}
		time.Sleep(100 * time.Millisecond)
	}
}
