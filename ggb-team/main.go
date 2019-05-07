package main

import (
	"flag"
	"log"

	"github.com/samuelvenzi/rcss-ggb/ggb-team/player"
)

const logPath = "/logs/ggb-team.log"

var verbose = flag.Bool("verbose", true, "print info level logs to stdout")

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	hostName := "rcssserver"

	player, err := player.NewPlayer("TeamGGB", hostName)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		err = player.SenseBody()
		if err != nil {
			log.Fatalln(err)
		}
		// time.Sleep(50 * time.Millisecond)
	}
}
