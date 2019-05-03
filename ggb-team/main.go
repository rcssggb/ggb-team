package main

import (
	"log"
	"net"
)

func main() {
	hostName := "rcssserver"
	portNum := "6000"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", RemoteAddr.String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	message := []byte("(init TeamGGB)\n")
	response := make([]byte, 1024)

	_, err = conn.WriteToUDP(message, RemoteAddr)

	if err != nil {
		log.Fatalln(err)
	}

	// receive message from server
	n, addr, err := conn.ReadFromUDP(response)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("UDP Server:", addr)
	log.Println("Received", n, "bytes from UDP server:", string(response))
}
