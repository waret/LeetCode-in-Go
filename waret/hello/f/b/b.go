package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	format = "2006-01-02 15:04:05.000"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	updateChan := make(chan int)

	go func(updateChan chan<- int) {
		for i := 0; i < 1; i++ {
			x := time.Duration(int64(time.Millisecond) * 100 * int64(rand.Uint64()%50+1))
			//fmt.Printf("[%v]: sleep for %v\n", time.Now().Format(format), x)
			time.Sleep(x)
			//fmt.Printf("[%v]: sleep for %v\n", time.Now().Format(format), x)
			updateChan <- i
		}
		time.Sleep(99999999999999)
	}(updateChan)

	go func(updateChan <-chan int) {
		minQuiet := time.Second * 5
		maxDelay := time.Second * 2
		var timeChan <-chan time.Time
		var startDebounce time.Time
		var lastConfigUpdateTime time.Time
		debouncedEvents := 0
		for {
			select {
			case <-updateChan:
				lastConfigUpdateTime = time.Now()
				if debouncedEvents == 0 {
					timeChan = time.After(minQuiet)
					startDebounce = lastConfigUpdateTime
				}
				debouncedEvents++
				fmt.Printf("[%v]: %v\n", time.Now().Format(format), debouncedEvents)
			case now := <-timeChan:
				fmt.Printf("[%v]: %v\n", time.Now().Format(format), now.Format(format))
				timeChan = nil
				eventDelay := now.Sub(startDebounce)
				quietTime := now.Sub(lastConfigUpdateTime)
				if eventDelay >= maxDelay || quietTime >= minQuiet {
					debouncedEvents = 0
					fmt.Printf("[%v]: process update\n", time.Now().Format(format))
					continue
				} else {
					fmt.Printf("[%v]: %v\n", time.Now().Format(format), minQuiet-quietTime)
					timeChan = time.After(minQuiet - quietTime)
					fmt.Printf("[%v]: %v\n", time.Now().Format(format), minQuiet-quietTime)
				}
			}
		}
		wg.Done()
	}(updateChan)

	wg.Wait()
}
