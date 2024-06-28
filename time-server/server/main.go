package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		buf := make([]byte, 2048)
		n, clientAddr, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))

		time := time.Now()
		_, err = ln.WriteToUDP([]byte(time.Format("15:04:05")), clientAddr)
		if err != nil {
			log.Println(err)
		}

	}
}
