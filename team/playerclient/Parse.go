package playerclient

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
		case ServerParamMsg:
			err = p.parseServerParam(m)
		case UnsupportedMsg:
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

	return nil
}

// parseServerParam parses (server_param ...
func (p *Player) parseServerParam(m Message) error {
	trimmedMsg := m.data
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(server_param ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")")

	for closeIdx := strings.Index(trimmedMsg, ")"); closeIdx != -1; closeIdx = strings.Index(trimmedMsg, ")") {
		currParam := trimmedMsg[1:closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+1 : len(trimmedMsg)-1]
		splitParam := strings.Split(currParam, " ")
		if len(splitParam) != 2 {
			return errors.New("something went wrong with server_param parsing")
		}
		paramName := splitParam[0]
		paramValString := splitParam[1]

		paramInterface, ok := p.serverParams[paramName]
		if !ok {
			log.Printf("unknown server param named %s\n", paramName)
			continue
		}

		switch param := paramInterface.(type) {
		case float64:
			paramFloat, err := strconv.ParseFloat(paramValString, 64)
			if err != nil {
				log.Printf("error parsing float server param: %s", err)
				continue
			}
			p.serverParams[paramName] = paramFloat
		case int:
			paramInt, err := strconv.ParseInt(paramValString, 10, 64)
			if err != nil {
				log.Printf("error parsing int server param: %s", err)
				continue
			}
			p.serverParams[paramName] = paramInt
		case bool:
			if paramValString == "1" {
				p.serverParams[paramName] = true
			} else if paramValString == "0" {
				p.serverParams[paramName] = false
			} else {
				log.Println("unknown bool format in server_param message")
				continue
			}
		case string:
			p.serverParams[paramName] = paramValString
		default:
			log.Printf("conversion not defined for param type %T, leaving default value %s\n", param, param)
		}
	}
	return nil
}

// parseSight parses (see 0 ((f r t) 55.7 3) ...
func (p *Player) parseSight(m Message) error {
	// TODO: implement sight parser
	trimmedMsg := m.data
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(see ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")")

	time := string(trimmedMsg[0])

	trimmedMsg = trimmedMsg[1:]

	for openIdx := strings.Index(trimmedMsg, "(("); openIdx != -1; openIdx = strings.Index(trimmedMsg, "((") {
		closeIdx := strings.Index(trimmedMsg, ")")
		objName := trimmedMsg[openIdx+2 : closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+2 : len(trimmedMsg)]

		closeIdx = strings.Index(trimmedMsg, ")")
		params := trimmedMsg[:closeIdx]
		trimmedMsg = trimmedMsg[closeIdx+1 : len(trimmedMsg)]

		splitParam := strings.Split(params, " ")

		_, ok := p.sightMap[objName]
		if !ok {
			log.Printf("unknown sight field named %s\n", objName)
			continue
		}
		timeInt, err := strconv.ParseInt(time, 10, 64)
		if err != nil {
			log.Printf("error parsing int time: %s", err)
		}
		p.sightMap[objName].time = timeInt

		switch nData := len(splitParam); nData {
		case 2:
			distance, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float distance: %s", err)
				continue
			}
			direction, err := strconv.ParseFloat(splitParam[1], 64)
			if err != nil {
				log.Printf("error parsing float direction: %s", err)
				continue
			}
			p.sightMap[objName].distance = distance
			p.sightMap[objName].direction = direction

		case 4:
			distance, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float distance: %s", err)
				continue
			}
			direction, err := strconv.ParseFloat(splitParam[1], 64)
			if err != nil {
				log.Printf("error parsing float direction: %s", err)
				continue
			}
			distChng, err := strconv.ParseFloat(splitParam[2], 64)
			if err != nil {
				log.Printf("error parsing float distance change: %s", err)
				continue
			}
			dirChng, err := strconv.ParseFloat(splitParam[3], 64)
			if err != nil {
				log.Printf("error parsing float direction change: %s", err)
				continue
			}
			p.sightMap[objName].distance = distance
			p.sightMap[objName].direction = direction
			p.sightMap[objName].distChng = distChng
			p.sightMap[objName].dirChng = dirChng
		case 6:
			distance, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float distance: %s", err)
				continue
			}
			direction, err := strconv.ParseFloat(splitParam[1], 64)
			if err != nil {
				log.Printf("error parsing float direction: %s", err)
				continue
			}
			distChng, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float distance change: %s", err)
				continue
			}
			dirChng, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float direction change: %s", err)
				continue
			}
			bodyDir, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float body direction: %s", err)
				continue
			}
			headDir, err := strconv.ParseFloat(splitParam[0], 64)
			if err != nil {
				log.Printf("error parsing float head direction: %s", err)
				continue
			}
			p.sightMap[objName].distance = distance
			p.sightMap[objName].direction = direction
			p.sightMap[objName].distChng = distChng
			p.sightMap[objName].dirChng = dirChng
			p.sightMap[objName].bodyDir = bodyDir
			p.sightMap[objName].headDir = headDir
		default:
			log.Printf("length of sight data for %s must be 2, 4 or 6", objName)
		}

	}

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
