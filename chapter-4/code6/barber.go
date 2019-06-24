package main

import (
	"fmt"
	"sync"
	//"time"
)

type Barber struct {
	ID         int
	ReadyToCut chan int
}

func (b *Barber) start(seatsMutex *sync.Mutex, customerReady chan int, seats *int) {
	fmt.Println("Barber", b.ID, "started...")
	for {
		fmt.Println("Waiting costumer to arrive or to ready to cut")
		costumerID := <-customerReady
		fmt.Println("Costumer", costumerID, "is ready")
		seatsMutex.Lock()
		*seats++
		b.ReadyToCut <- b.ID
		fmt.Println("Will now start cutting")
		seatsMutex.Unlock()
		// CUT HAIR HERE
		// time.Sleep(time.Millisecond * 100)
	}

}
