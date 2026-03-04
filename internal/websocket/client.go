package websocket

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Client represents a single WebSocket client.
type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	mu   sync.Mutex
}

// NewClient creates a new Client instance.
func NewClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
}

// ReadMessages reads messages from the WebSocket connection.
func (c *Client) ReadMessages() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		c.hub.broadcast <- message
	}
}

// WriteMessages writes messages to the WebSocket connection.
func (c *Client) WriteMessages() {
	defer c.conn.Close()
	for message := range c.send {
		c.mu.Lock()
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		c.mu.Unlock()
		if err != nil {
			break
		}
	}
}
