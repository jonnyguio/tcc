package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"tcc-chapter4-readerwriter-code3/reader"
	"tcc-chapter4-readerwriter-code3/writer"
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

	streamSize := 32
	stream := make([]byte, streamSize)
	readers := &sync.Mutex{}
	writers := &sync.Mutex{}

	for i := 0; i < numWriters; i++ {
		writer := &writer.Writer{ID: i + 1, Readers: readers, Writers: writers}
		go writer.Start(stream, streamSize)
	}
	for i := 0; i < numReaders; i++ {
		reader := &reader.Reader{ID: i + 1, Readers: readers, Writers: writers}
		go reader.Start(stream, streamSize)
	}
	time.Sleep(100 * time.Second)
}
