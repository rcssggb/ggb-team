package player

import (
	"log"
	"time"
)

// Listen continuously receives and parses server messages and updates the
// Player object accordingly
func (p *Player) Listen() {
	response := make([]byte, 8192)

	for {
		p.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		_, _, err := p.conn.ReadFromUDP(response)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(string(response))
	}
}
