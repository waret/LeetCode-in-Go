package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// WaitSignal awaits for SIGINT or SIGTERM and closes the channel
func WaitSignal(stop chan struct{}) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("stop")
	close(stop)
}

func main() {
	stop := make(chan struct{})
	time.After(time.Second)
	go func(stop <-chan struct{}) {
		tick := time.NewTicker(time.Second)
		for {
			select {
			case <-stop:
				fmt.Println("stop")
				return
			case <-tick.C:
				fmt.Println("hello")
			}
		}
	}(stop)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("i = %v\n", i)
			time.Sleep(time.Millisecond * 1500)
		}
	}()

	WaitSignal(stop)
}
