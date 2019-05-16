package main

import (
	"fmt"
	"sync"
	"net/http"
)

func main() {
	downloadSuccess := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if <-downloadSuccess == true {
			fmt.Println("Download success")
		} else {
			fmt.Println("Download failed")
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := http.Get("http://example.com")
		if err != nil {
			downloadSuccess <- false
			return
		}
		downloadSuccess <- true
	}()
	wg.Wait()
}
