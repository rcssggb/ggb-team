package playerclient

// MessageType is the type for message type constants
type MessageType byte

const (
	// DisabledMsg is the type for all unhandled packets
	DisabledMsg MessageType = iota + 0

	// ErrorMsg is the type for all error messages received
	ErrorMsg

	// InitMsg is the type for the init message received when connecting
	InitMsg

	// SightMsg is the type for visual sensor messages
	SightMsg

	// BodyMsg is the type for body sensor messages
	BodyMsg

	// PlayerTypeMsg is the type for heterogenous player message
	PlayerTypeMsg
)
