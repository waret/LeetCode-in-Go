package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	for i := 0; i < 1000; i++ {
	}
	elapsed := time.Since(start)

	fmt.Println(elapsed)
	fmt.Println("xxxxxxxxx %s", time.Now().String())

	aa := time.Duration(int64(5 * 1e9))

	select {
	case <-time.After(aa):
	}
	fmt.Println(aa)

	fmt.Println("xxxxxxxxx %s", time.Now().String())

	select {
	case <-time.After(elapsed - aa):
	}

	fmt.Println(elapsed-aa)

	fmt.Println("xxxxxxxxx %s", time.Now().String())
}
