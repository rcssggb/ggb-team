package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	hostName := "udp-echo"
	portNum := "6000"

	service := hostName + ":" + portNum

	RemoteAddr, err := net.ResolveUDPAddr("udp", service)

	//LocalAddr := nil
	// see https://golang.org/pkg/net/#DialUDP

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	// note : you can use net.ResolveUDPAddr for LocalAddr as well
	//        for this tutorial simplicity sake, we will just use nil

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	defer conn.Close()

	// write a message to server
	message := []byte("(init TeamGGB)\n")
	response := make([]byte, 1024)

	_, err = conn.Write(message)

	if err != nil {
		log.Fatalln(err)
	}

	// receive message from server
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	n, addr, err := conn.ReadFromUDP(response)
	// n, err := conn.Read(response)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("UDP Server: ", addr)
	fmt.Println("Received", n, "bytes from UDP server: ", string(response))
}
