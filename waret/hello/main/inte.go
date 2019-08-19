package main

import "fmt"

type Limit interface {
    Func01(para int) int
    Func02(para int) int
}

type Limiter struct {
	connectionBurstSize int
}

func (c *Limiter) Func02(para int) int {
    fmt.Printf("connectionBurstSize: %d ", c.connectionBurstSize)
    fmt.Printf("para: %d ", para)
    return para
}

func (c *Limiter) Func01(para int) int {
    fmt.Printf("connectionBurstSize: %d ", c.connectionBurstSize)
    fmt.Printf("para: %d ", para)
    return para
}

func inner(limit Limit) {
    limit.Func01(1)
    limit.Func02(2)
}

func main() {
    var limit Limit
    limit = &Limiter{
        connectionBurstSize: 10,
    }
    inner(limit)
}
