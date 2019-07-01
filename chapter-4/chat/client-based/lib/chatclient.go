package lib

import "net"

type ChatClient struct {
	Name     string
	Listener net.Listener
}
