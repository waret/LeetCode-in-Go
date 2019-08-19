package main
import "fmt"

func getSequence() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    fmt.Printf("%d ", getSequence()())
    fmt.Printf("%d ", getSequence()())
    fmt.Printf("%d ", getSequence()())

    nextNumber := getSequence()
    fmt.Printf("%d ", nextNumber())
    fmt.Printf("%d ", nextNumber())
    fmt.Printf("%d ", nextNumber())
}
