package main

import (
	"fmt"
	"sync"
)

type Costumer struct {
	ID            int
	CostumerReady chan int
}

func (c Costumer) start(seatsMutex *sync.Mutex, barberReady chan int, seats *int) {
	fmt.Println("Costumer", c.ID, "started...")
	seatsMutex.Lock()
	if *seats > 0 {
		*seats--
		seatsMutex.Unlock()
		fmt.Println("waiting barber to be ready or to receive my arrival")
		c.CostumerReady <- c.ID
		fmt.Println("barber knows i have arrived")
		barberID := <-barberReady
		fmt.Println("i will now have my hair cut by barber", barberID)
		// HAVE HAIR CUT HERE
		// time.Sleep(time.Millisecond * 100)
	} else {
		seatsMutex.Unlock()
		fmt.Println("No place to wait, going away...")
	}
}
