package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Producer struct {
	ID         int
	stream     chan []byte
	streamSize int
}

func (p *Producer) start(stream chan []byte, streamSize int, wg *sync.WaitGroup) {
	p.stream = stream
	p.streamSize = streamSize

	randomInput := make([]byte, streamSize)
	for {
		n, err := rand.Read(randomInput)
		if err != nil || n != p.streamSize {
			fmt.Println("Error creating randomInput, exiting producer.")
			break
		}
		fmt.Printf("Producer %d writing randomInput\n", p.ID)
		p.stream <- randomInput
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}
