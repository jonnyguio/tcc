package main

import (
	"errors"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"../lib"
)

var clients = &sync.Map{}

type Server struct{}

func (s *Server) Register(name interface{}, cc *lib.ChatClient) error {
	_, ok := clients.Load(name)
	if ok {
		return errors.New("Name in use")
	}
	clients.Store(name, cc)
	return nil
}

func (s *Server) Connect(name interface{}, fn *func() net.Conn) error {
	client, ok := clients.Load(name)
	cclient, ok := client.(*lib.ChatClient)
	if ok {
		var err error
		var conn net.Conn
		*fn = func() net.Conn {
			conn, err = net.Dial(cclient.Listener.Addr().Network(), cclient.Listener.Addr().String())
			return conn
		}
		return err
	} else {
		*fn = func() net.Conn {
			return nil
		}
		return errors.New("Failed to find client")
	}
	return nil
}

func (s *Server) Deregister(name interface{}, cc *lib.ChatClient) error {
	clients.Delete(name)
	return nil
}

func main() {
	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	http.Serve(listener, nil)

	/*clients := make(map[string]net.Conn)
	dead := make(chan string)
	messages := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				panic(err)
			}
			_, err = conn.Write([]byte("Welcome to the chat!\n"))
			if err != nil {
				break
			}
			reader := bufio.NewReader(conn)
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
	}*/
}
