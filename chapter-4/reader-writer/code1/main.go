package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <number of readers> <number of writers>\n", os.Args[0])
		os.Exit(1)
	}
	numReaders, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Failed to convert consumer argument as number")
		panic(err)
	}
	numWriters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Failed to convert producer argument as number")
		panic(err)
	}

	stream := make([]byte, 16)
	streamSize := 16
	mutex := sync.RWMutex{}

	for i := 0; i < numWriters; i++ {
		writer := Writer{ID: i + 1, Mutex: &mutex}
		go writer.start(stream, streamSize)
	}
	for i := 0; i < numReaders; i++ {
		reader := &Reader{ID: i + 1, Mutex: &mutex}
		go reader.start(stream, streamSize)
	}
	time.Sleep(4 * time.Second)
}
