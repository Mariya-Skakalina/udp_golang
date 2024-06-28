package main

import (
	"log"
	"net"
	"strings"
	"unicode/utf8"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	buf := make([]byte, 1024)

	for {
		n, clientAddr, err := ln.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}
		text := string(buf[:n])
		msg := strings.TrimSpace(text)
		if utf8.RuneCountInString(msg) >= 4 {
			anw := "Ваша строка длинее 4 символов"
			ln.WriteToUDP([]byte(anw), clientAddr)
		} else {
			anw := "You are LOX"
			ln.WriteToUDP([]byte(anw), clientAddr)
		}

	}
}
