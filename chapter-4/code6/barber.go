package main

import (
	"fmt"
	"time"
	//"time"
)

type Barber struct {
	ID         int
	ReadyToCut chan int
}

func (b *Barber) start(customerReady chan int, seatsCh chan int) {
	fmt.Println("Barber", b.ID, "started...")
	for {
		fmt.Println("Waiting costumer to arrive or to ready to cut")
		costumerID := <-customerReady
		fmt.Println("Costumer", costumerID, "is ready")
		seatsCh <- <-seatsCh + 1
		b.ReadyToCut <- b.ID
		fmt.Println("Will now start cutting")
		// CUT HAIR HERE
		time.Sleep(time.Millisecond * 3000)
	}

}
