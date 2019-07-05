package main

import (
	"math/rand"
	"time"
)

func main() {
	costumerChannel := make(chan int)
	costumers := []Costumer{}
	for i := 0; i < 10; i++ {
		costumers = append(costumers, Costumer{
			ID:            i + 1,
			CostumerReady: costumerChannel,
		})
	}
	barber := &Barber{ID: 1, ReadyToCut: make(chan int)}

	seatsCh := make(chan int, 1)
	seatsCh <- 5
	go barber.start(costumerChannel, seatsCh)
	for i := 0; i < 10; i++ {
		turn := rand.Int() % 10
		go costumers[turn].start(barber.ReadyToCut, seatsCh)
	}
	time.Sleep(30 * time.Second)
}
