package player

import (
	"log"
	"time"
)

// SenseBody fetches Body Sensor data for the player
func (p *Player) SenseBody() error {
	// Send request
	_, err := p.conn.Write([]byte("(sense_body)"))
	if err != nil {
		log.Println(err)
		return err
	}

	response := make([]byte, 1024)

	p.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, _, err = p.conn.ReadFromUDP(response)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(string(response))

	return nil
}
