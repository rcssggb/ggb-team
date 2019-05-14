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

// Type parses and returns the MessageType for the message
func (m Message) Type() MessageType {
	switch rune(m.data[1]) {
	case 'i': // ( i nit ...
		return InitMsg
	case 'e': // ( e rror ...
		return ErrorMsg
	default:
		return DisabledMsg
	}
}
