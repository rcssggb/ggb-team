package main

import (
	"fmt"
	"net"
	"os"
)

// CheckError checks for errors
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":6174")
	CheckError(err)
	fmt.Println("listening on :6174")

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Printf("received: %s from: %s\n", string(buf[0:n]), addr)

		if err != nil {
			fmt.Println("error: ", err)
		}

		ServerConn.WriteTo(buf[0:n], addr)
	}
}
