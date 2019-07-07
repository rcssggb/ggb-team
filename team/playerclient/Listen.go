package playerclient

import (
	"bytes"
	"log"
	"time"
)

// Listen continuously receives and posts server messages to cmdChannel
func (p *Player) Listen() {
	response := make([]byte, 8192)

	for {
		p.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		_, _, err := p.conn.ReadFromUDP(response)
		now := time.Now()
		if err != nil {
			log.Println(err)
			continue
		}
		data := bytes.Trim(response, "\x00")
		p.cmdChannel <- Message{
			timestamp: now,
			data:      string(data),
		}
	}
}
