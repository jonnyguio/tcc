package main

import (
	"sync"
	"time"
	"math/rand"
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

	mutex := &sync.Mutex{}
	seats := new(int)
	*seats = 5
	go barber.start(mutex, costumerChannel, seats)
	for i := 0; i < 100; i++ {
		turn := rand.Int() % 10
		// fmt.Println(costumers[turn].ID)
		go costumers[turn].start(mutex, barber.ReadyToCut, seats)
	}
	time.Sleep(30 * time.Second)
}
