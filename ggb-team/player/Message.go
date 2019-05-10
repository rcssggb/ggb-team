package player

import "time"

// Message is the struct used to move and parse RCSS commands
type Message struct {
	timestamp time.Time
	data      []byte
}

// String implements the native Stringer interface
func (m Message) String() string {
	return string(m.data)
}
