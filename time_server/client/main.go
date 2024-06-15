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
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 2048)
		var msg string
		fmt.Print("Enter message: ")
		fmt.Scanln(&msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
