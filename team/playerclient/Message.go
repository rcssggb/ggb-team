package playerclient

import (
	"strings"
	"time"
)

// Message is the struct used to move and parse RCSS commands
type Message struct {
	timestamp time.Time
	data      string
}

// String implements the native Stringer interface
func (m Message) String() string {
	return string(m.data)
}

// Type parses and returns the MessageType for the message
func (m Message) Type() (mType MessageType) {
	switch {
	case strings.HasPrefix(m.data, "(see "):
		mType = SightMsg
	case strings.HasPrefix(m.data, "(sense_body "):
		mType = BodyMsg
	case strings.HasPrefix(m.data, "(error "):
		mType = ErrorMsg
	case strings.HasPrefix(m.data, "(init "):
		mType = InitMsg
	case strings.HasPrefix(m.data, "(player_type "):
		mType = PlayerTypeMsg
	case strings.HasPrefix(m.data, "(player_param "):
		mType = PlayerParamMsg
	case strings.HasPrefix(m.data, "(server_param "):
		mType = ServerParamMsg
	// Unsupported msgs below
	case strings.HasPrefix(m.data, "(score "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(warning "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(ok "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(change_player_type "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(fullstate "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(hear "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(change "):
		mType = UnsupportedMsg
	case strings.HasPrefix(m.data, "(clang "):
		mType = UnsupportedMsg
	default:
		mType = UnsupportedMsg
	}
	return mType
}
