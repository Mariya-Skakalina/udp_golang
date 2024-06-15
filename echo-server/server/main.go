package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Listening server port 8080")

	for {
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}

		msg := fmt.Sprintln("Your message: ", string(buf[:n]))
		_, err = conn.WriteToUDP([]byte(msg), clientAddr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
