package main

import (
	"fmt"
	"net"
	"os"

	"../lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ./%s <name> <name_to_connect>\n")
	}
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	chatClient := &lib.ChatClient{
		os.Args[1],
		listener,
	}
	if len(os.Args > 2)
}
