package main

import (
	"fmt"
	"net"
)

func echoResponse(conn *net.UDPConn, addr *net.UDPAddr, msg string) {
	_, err := conn.WriteToUDP([]byte(msg), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v\n", err)
	}
}

func main() {
	port := 6000
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP("127.0.0.1"),
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	fmt.Printf("Listening on port %d\n", port)
	for {
		_, remoteaddr, err := server.ReadFromUDP(p)
		msg := string(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go echoResponse(server, remoteaddr, msg)
	}
}
