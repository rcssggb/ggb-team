package playerclient

import (
	"fmt"
	"log"
)

// Parse continuously receives messages from cmdChannel and parses them to update the player object
func (p *Player) Parse() {
	var m Message
	var err error
	for {
		m = <-p.cmdChannel
		switch m.Type() {
		case InitMsg:
			err = p.parseInit(m)
		case ErrorMsg:
			err = p.parseError(m)
		case SightMsg:
			err = p.parseSight(m)
		case BodyMsg:
			err = p.parseBody(m)
		case PlayerTypeMsg:
			err = p.parsePlayerType(m)
		case UnsupportedMsg:
			if len(m.data) > 64 {
				fmt.Printf("Ignoring %s...\n", string(m.data[0:64]))
			} else {
				fmt.Printf("Ignoring %s\n", string(m.data))
			}
			continue
		}
		if err != nil {
			fmt.Println(err)
			err = nil
		}
	}
}

// parseInit parses (init Side Unum PlayMode)
func (p *Player) parseInit(m Message) error {
	var side rune
	var unum int
	var playMode string

	_, err := fmt.Sscanf(m.data, "(init %c %d %s", &side, &unum, &playMode)
	if err != nil {
		return err
	}
	playMode = playMode[0 : len(playMode)-1] // trim out last )

	p.teamSide = (side == 'r')
	p.shirtNum = unum
	p.playMode = playMode

	fmt.Println(m, p.teamSide, p.shirtNum, p.playMode)

	return nil
}

// parseSight parses (see 0 ((f r t) 55.7 3) ...
func (p *Player) parseSight(m Message) error {
	// TODO: implement sight parser
	return nil
}

// parseBody parses (sense_body 0 (view_mode high normal) ...
func (p *Player) parseBody(m Message) error {
	// TODO: implement body parser
	return nil
}

// parsePlayerType parses (player_type (id 17)(player_speed_max ...
func (p *Player) parsePlayerType(m Message) error {
	// TODO: implement player type parser
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
