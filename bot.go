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
	url               = "https://camo.githubusercontent.com/ccabda83743448d75055da131db332c2d62620cf73b695bdcb67dd00a7423b43/68747470733a2f2f6170692e76697369746f7262616467652e696f2f6170692f56697369746f724869743f757365723d786361746265617226636f756e74436f6c6f72636f756e74436f6c6f7226636f756e74436f6c6f723d253233303036454646"
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
