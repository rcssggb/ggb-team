package playerclient

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
func (m Message) Type() (mType MessageType) {
	switch rune(m.data[1]) {
	case 'i': // ( i nit ...
		mType = InitMsg
	case 's': // (s
		switch rune(m.data[3]) {
		case 'e':
			switch rune(m.data[4]) {
			case ' ': // (s e e ' '
				mType = SightMsg
			}
		case 'n': // (s e n se_body ...
			mType = BodyMsg
		}
	case 'e': // ( e rror ...
		mType = ErrorMsg
	case 'p': // ( p layer ...
		switch rune(m.data[8]) {
		case 't': // (player- t ype
			mType = PlayerTypeMsg
		}
	default:
		mType = DisabledMsg
	}
	return mType
}
