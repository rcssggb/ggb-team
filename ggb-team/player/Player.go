package player

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/samuelvenzi/rcss-ggb/ggb-team/rcssparser"
)

// Player ...
type Player struct {
	conn     *net.UDPConn
	teamSide string
	shirtNum int
	playMode string
}

// IPlayer ...
type IPlayer interface {
	SenseBody() error
}

// NewPlayer is the constructor for the Player struct
func NewPlayer(teamName, serverIP string) (IPlayer, error) {
	// Open listener socker to request player connection
	serverHost := serverIP + ":6000"
	serverAddr, err := net.ResolveUDPAddr("udp", serverHost)
	if err != nil {
		return nil, err
	}
	tmpConn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, err
	}

	// Send connect message
	cmdMessage := fmt.Sprintf("(init %s)", teamName)
	_, err = tmpConn.WriteToUDP([]byte(cmdMessage), serverAddr)
	if err != nil {
		return nil, err
	}

	// Receive response
	response := make([]byte, 1024)
	tmpConn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, addr, err := tmpConn.ReadFromUDP(response)
	if err != nil {
		return nil, err
	}
	tmpConn.Close()

	// Parse response
	initResponse, err := rcssparser.ParseInit(response)
	if err != nil {
		return nil, err
	}

	// Dial to dedicated player port
	playerPort := addr.Port
	playerHost := serverIP + ":" + strconv.Itoa(playerPort)
	playerAddr, err := net.ResolveUDPAddr("udp", playerHost)
	if err != nil {
		return nil, err
	}

	// Instantiate new Player struct
	newPlayer := &Player{}

	// Init newPlayer struct
	conn, err := net.DialUDP("udp", nil, playerAddr)
	newPlayer.conn = conn

	// (init Side Unum PlayMode)
	newPlayer.teamSide = initResponse[1]
	shirtNum, err := strconv.Atoi(initResponse[2])
	if err != nil {
		return nil, err
	}
	newPlayer.shirtNum = shirtNum
	newPlayer.playMode = initResponse[3]

	return newPlayer, nil
}
