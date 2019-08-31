package main

import "fmt"

type I interface {
	walk()
}
type J interface {
	quack()
}
type K interface {
	bark()
}
type S struct{}

func (s S) walk()  {}
func (s S) quack() {}
func main() {
	var i interface{}
	i = S{}
	v1, ok1 := i.(J)
	fmt.Printf("%T, %v\n", v1, ok1)
	v2, ok2 := i.(K)
	fmt.Printf("%T, %v\n", v2, ok2) // panic: interface conversion: main.S is not main.K: missing method bark
}
