package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
)

func main() {
	_, err := os.Stat("upload")
	if os.IsNotExist(err) {
		err := os.Mkdir("upload", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	id := uuid.New()
	filePath := fmt.Sprintf("upload/%s.txt", id.String())
	file, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, _, err := ln.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		if n == 0 {
			break
		}
		if _, err = file.Write(buf[:n]); err != nil {
			log.Println(err)
			continue
		}
	}
	fmt.Println("File received successfully")
}
