package rcssparser

import "fmt"

// ParseInit parses (init Side Unum PlayMode)
func ParseInit(response []byte) ([]string, error) {
	switch string(response) {
	case "(error no_more_team_or_player)":
		return nil, fmt.Errorf("Unable to create new team OR team already full")
	case "(error reconnect)":
		return nil, fmt.Errorf("Player already connected")
	}
	return breakArgs(string(response)), nil
}
