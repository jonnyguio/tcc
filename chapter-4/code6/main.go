package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	costumerChannel := make(chan int)
	costumers := []Costumer{}
	for i := 0; i < 5; i++ {
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
	for _, costumer := range costumers {
		fmt.Println(costumer.ID)
		go costumer.start(mutex, barber.ReadyToCut, seats)
	}
	time.Sleep(10 * time.Second)
}
