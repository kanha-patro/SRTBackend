package websocket

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func NewMessage(messageType string, payload interface{}) *Message {
	return &Message{
		Type:    messageType,
		Payload: payload,
	}
}
