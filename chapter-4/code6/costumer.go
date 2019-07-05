package main

import (
	"fmt"
	"time"
)

type Costumer struct {
	ID            int
	CostumerReady chan int
}

func (c Costumer) start(barberReady chan int, seatsCh chan int) {
	fmt.Println("Costumer", c.ID, "started...")
	seats := <-seatsCh
	if seats > 0 {
		seatsCh <- seats - 1
		fmt.Println("waiting barber to be ready or to receive my arrival")
		c.CostumerReady <- c.ID
		fmt.Println("barber knows i have arrived")
		barberID := <-barberReady
		fmt.Println("i will now have my hair cut by barber", barberID)
		// HAVE HAIR CUT HERE
		time.Sleep(time.Millisecond * 3000)
	} else {
		seatsCh <- seats
		fmt.Println("No place to wait, going away...")
	}
}
