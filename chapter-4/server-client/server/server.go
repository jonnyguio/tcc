package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Can't listen to socket")
			break
		}
		go func() {
			message, _ := bufio.NewReader(conn).ReadString('\n')

			fmt.Print("Message received:", string(message))

			newMessage := strings.ToUpper(message)

			conn.Write([]byte(newMessage + "\n"))
		}()
	}
}
