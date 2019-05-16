package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <number of consumers> <number of producers>\n", os.Args[0])
		os.Exit(1)
	}
	consumers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert consumer argument as number")
		panic(err)
	}
	producers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Failed to convert producer argument as number")
		panic(err)
	}
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UTC().UnixNano())
	streamSize := 32
	stream := make(chan []byte)

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		go consumer(stream, streamSize, wg)
	}
	for i := 0; i < producers; i++ {
		wg.Add(1)
		go producer(stream, streamSize, wg)
	}
	wg.Wait()
}
