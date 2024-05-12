package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var totalSuccessfulRequests int64

func makeRequests(url string, requestsPerSecond int) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
		} else {
			atomic.AddInt64(&totalSuccessfulRequests, 1)
			resp.Body.Close() 
		}
		time.Sleep(time.Second / time.Duration(requestsPerSecond))
	}
}

func printProgress() {
	for {
		successful := atomic.LoadInt64(&totalSuccessfulRequests)
		fmt.Printf("Total successful requests made: %d\n", successful)
		time.Sleep(time.Second)
	}
}

func main() {
	var url string
	var requestsPerSecond int

	fmt.Print("Enter the URL: ")
	fmt.Scanln(&url)

	fmt.Print("Enter the amount of Threads: ")
	fmt.Scanln(&requestsPerSecond)

	var wg sync.WaitGroup
	for i := 0; i < requestsPerSecond; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequests(url, requestsPerSecond)
		}()
	}

	go printProgress()
	wg.Wait()
}
