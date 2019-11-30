package main

import (
	"flag"
	"log"
	"time"

	playerclient "github.com/rcssggb/ggb-lib/playerclient"
)

const logPath = "/logs/ggb-team.log"

var verbose = flag.Bool("verbose", true, "print info level logs to stdout")

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	hostName := "rcssserver"

	player, err := playerclient.NewPlayerClient("TeamGGB", hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	for {
		currentTime := player.Time()

		if currentTime == 0 {
			player.Move(currentTime, -0.1, 0)
		} else {
			if currentTime%2 == 0 {
				player.Dash(currentTime, 100)
				player.TurnNeck(currentTime, 45)
			} else {
				player.Turn(currentTime, 10)
				player.TurnNeck(currentTime, -45)
			}
		}
		err = player.Error()
		if err != nil {
			log.Println(err)
		}

		for player.Time() <= currentTime {
			time.Sleep(10 * time.Millisecond)
		}
	}
}
