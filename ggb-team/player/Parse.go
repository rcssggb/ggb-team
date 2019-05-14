package player

import (
	"fmt"
	"log"
	"time"
)

// Parse continuously receives messages from cmdChannel and parses them to update the player object
func (p *Player) Parse() {
	var m Message
	for {
		m = <-p.cmdChannel
		switch m.Type() {
		case InitMsg:
			p.parseInit(m)
		case ErrorMsg:
			p.parseError(m)
		case DisabledMsg:
			if len(m.data) > 64 {
				fmt.Printf("Ignoring %s...\n", string(m.data[0:64]))
			} else {
				fmt.Printf("Ignoring %s\n", string(m.data))
			}
			continue
		}
		fmt.Printf("Delay: %s\n", time.Now().Sub(m.timestamp))
	}
}

// parseInit parses (init Side Unum PlayMode)
func (p *Player) parseInit(m Message) error {
	var side rune
	var unum int
	var playMode string

	//TODO: trim ) properly with scan format
	_, err := fmt.Sscanf(string(m.data), "(init %c %d %s)", &side, &unum, &playMode)
	if err != nil {
		return err
	}
	playMode = playMode[0 : len(playMode)-2] // trim out last )

	p.teamSide = side
	p.shirtNum = unum
	p.playMode = playMode

	return nil
}

func (p *Player) parseError(m Message) error {
	var errMsg string
	switch rune(m.data[7]) {
	case 'n': // (error n o_more_team_or_player)
		errMsg = fmt.Sprintf("Team %s is full or unable to create third team\n", p.teamName)
		log.Fatal(errMsg)
	case 'r': // (error r econnect)
		errMsg = fmt.Sprintf("Player already connected")
		log.Fatal(errMsg)
	}
	return fmt.Errorf(errMsg)
}
