package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const (
	requestsPerSecond = 500
	url               = "URL TO YOUR VIEWS COUNTER ON GITHUB"
)

var totalSuccessfulRequests int64

func makeRequests() {
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
		} else {
			atomic.AddInt64(&totalSuccessfulRequests, 1)
			resp.Body.Close()
		}
		time.Sleep(time.Second / requestsPerSecond)
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
	var wg sync.WaitGroup
	for i := 0; i < requestsPerSecond; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequests()
		}()
	}

	go printProgress()
	wg.Wait()
}
