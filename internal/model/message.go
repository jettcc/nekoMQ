package model

type Message struct {
	// data is the message content
	// The first 16 bits, as uuid, are unique identifiers
	data []byte
}

func NewMessage(data []byte) *Message {
	return &Message{data: data}
}

func (m *Message) Data() []byte {
	return m.data
}

func (m *Message) MsgUuid() []byte {
	return m.data[:16]
}

func (m *Message) MsgContent() []byte {
	return m.data[16:]
}
