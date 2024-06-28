package main

import (
	"fmt"
	"net"
)

func main() {
	broadcastIP := "255.255.255.255"
	port := "12345"
	message := []byte("Hello, broadcast!")

	addr, err := net.ResolveUDPAddr("udp", broadcastIP+":"+port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent to broadcast address:", broadcastIP)
}
