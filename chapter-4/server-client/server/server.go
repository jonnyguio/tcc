package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("Message received:", string(message))

		newMessage := strings.ToUpper(message)

		conn.Write([]byte(newMessage + "\n"))
	}
}
