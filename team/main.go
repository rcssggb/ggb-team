package main

import (
	"flag"
	"log"
	"time"

	"./playerclient"
)

const logPath = "/logs/ggb-team.log"

var verbose = flag.Bool("verbose", true, "print info level logs to stdout")

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	hostName := "rcssserver"

	_, err := playerclient.NewPlayer("TeamGGB", hostName)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)
	return
}
