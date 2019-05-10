package player

import (
	"fmt"
	"time"
)

// Parse continuously receives messages from cmdChannel and parses them to update the player object
func (p *Player) Parse() {
	var m Message
	for {
		m = <-p.cmdChannel
		fmt.Printf("Delay: %s\n", time.Now().Sub(m.timestamp))
	}
}
