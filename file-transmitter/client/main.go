package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	var text string
	fmt.Println("Выберите файл формата txt")
	fmt.Scan(&text)
	filePath := strings.TrimSpace(text)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == os.ErrClosed || err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if _, err := conn.Write(buffer[:n]); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("File sent successfully")
}
