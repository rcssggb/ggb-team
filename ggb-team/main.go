package main

import (
	"flag"
	"log"
	"time"

	"github.com/samuelvenzi/rcss-ggb/ggb-team/player"
)

const logPath = "/logs/ggb-team.log"

var verbose = flag.Bool("verbose", true, "print info level logs to stdout")

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	hostName := "rcssserver"

	_, err := player.NewPlayer("TeamGGB", hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)
	return
}
