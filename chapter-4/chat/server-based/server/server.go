package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	clients := make(map[string]net.Conn)
	dead := make(chan string)
	messages := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				panic(err)
			}
			reader := bufio.NewReader(conn)
			_, err = conn.Write([]byte("Welcome to the chat! Please, type your name!\n"))
			if err != nil {
				break
			}
			name, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			name = name[:len(name)-2]
			if _, ok := clients[name]; ok {
				for {
					temp := name + strconv.Itoa(int(rand.Int31n(10e07)))
					if _, ok := clients[temp]; ok {
						continue
					}
					name = temp
					break
				}
			}
			clients[name] = conn
			log.Printf("New client: %s! There are %d clients on the chat\n", name, len(clients))

			go func(conn net.Conn, name string) {
				messages <- fmt.Sprintf("O usuario %s entrou do chat.\n", name)
				for {
					message, err := reader.ReadString('\n')
					if err != nil {
						break
					}
					messages <- fmt.Sprintf("[%s]: %s", name, message)
				}
				dead <- name
			}(conn, name)
		}
	}()

	for {
		select {
		case message := <-messages:
			for name, conn := range clients {
				go func(conn net.Conn) {
					_, err := conn.Write([]byte(message))
					if err != nil {
						dead <- name
					}
				}(conn)
			}
			log.Printf("New message: \"%s\". Broadcast to %d clients\n", message[:len(message)-2], len(clients))

		case name := <-dead:
			go func() {
				log.Printf("Client %s disconnected\n", name)
				delete(clients, name)
				messages <- fmt.Sprintf("O usuario %s saiu do chat.\n", name)
			}()
		}
	}
}
