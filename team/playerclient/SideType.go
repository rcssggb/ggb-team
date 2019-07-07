package playerclient

// SideType is an alias for defining the field side where team is playing
type SideType bool

const (
	// LeftSide ...
	LeftSide SideType = false
	// RightSide ...
	RightSide SideType = true
)

func (s SideType) String() string {
	if s {
		return "right"
	} else {
		return "left"
	}
}
