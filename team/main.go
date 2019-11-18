package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/rcssggb/ggb-lib/common/field"
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
		visualSensor := player.See()
		log.Printf("%+v\n", visualSensor)

		posX := field.MaxX * -rand.Float64()
		posY := field.MaxY * (2*rand.Float64() - 1)

		player.Move(visualSensor.Time, posX, posY)
		err = player.Error()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
