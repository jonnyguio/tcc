package main

import (
	"fmt"
	"net/http"
	"time"
)

func downloadPage(url string, page chan string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bytes := make([]byte, 1024)
	resp.Body.Read(bytes)
	page <- string(bytes)
}

func main() {
	page := make(chan string)
	urls := []string{"http://example.com", "http://ufrj.br/index.html", "http://github.com/index.html"}
	currentURL := 0
	for {
		select {
		case content := <-page:
			fmt.Println(content)
		case <-time.After(time.Second * 1):
			if currentURL < len(urls) {
				url := urls[currentURL]
				go downloadPage(url, page)
				fmt.Println("Getting next page")
				currentURL++
			} else {
				fmt.Println("No more pages to download, exiting.")
				return
			}
		}
	}

}
