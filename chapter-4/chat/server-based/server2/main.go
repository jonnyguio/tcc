package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	clients := &sync.Map{}
	dead := make(chan string)
	messages := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())
	go receivePeople(server, clients, messages, dead)
	for {
		select {
		case message := <-messages:
			clients.Range(func(name, conn interface{}) bool {
				go func(conn net.Conn) {
					_, err := conn.Write([]byte(message))
					if err != nil {
						dead <- name.(string)
					}
				}(conn.(net.Conn))
				return true
			})
			log.Printf("New message: \"%s\"\n", message[:len(message)-2])
		case name := <-dead:
			go func() {
				log.Printf("Client %s disconnected\n", name)
				clients.Delete(name)
				messages <- fmt.Sprintf("O usuario %s saiu do chat.\n", name)
			}()
		}
	}
}

func receivePeople(server net.Listener, clients *sync.Map, messages chan string, dead chan string) {
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)
			_, err = conn.Write([]byte("Welcome to the chat! Please, type your name!\n"))
			if err != nil {
				return
			}
			name, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			name = name[:len(name)-2]
			if _, ok := clients.Load(name); ok {
				for {
					temp := name + strconv.Itoa(int(rand.Int31n(10e07)))
					if _, ok = clients.Load(temp); ok {
						continue
					}
					name = temp
					break
				}
			}
			clients.Store(name, conn)
			log.Printf("New client: %s!\n", name)

			messages <- fmt.Sprintf("O usuario %s entrou do chat.\n", name)
			for {
				message, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				messages <- fmt.Sprintf("[%s]: %s", name, message)
			}
			dead <- name
		}(conn)
	}
}
