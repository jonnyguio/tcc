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
		fmt.Printf("Usage: %s <number of consumers> <number of producers> <buffer size>\n", os.Args[0])
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
	stream := make(chan []byte, producers)
	streamSize := 32

	for i := 0; i < consumers; i++ {
		wg.Add(1)
		newConsumer := Consumer{ID: i + 1}
		go newConsumer.start(stream, streamSize, &wg)
	}
	for i := 0; i < producers; i++ {
		wg.Add(1)
		newProducer := Producer{ID: i + 1}
		go newProducer.start(stream, streamSize, &wg)
	}

	wg.Wait()
}
