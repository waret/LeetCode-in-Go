package main

import "fmt"

type I interface {
	walk()
}
type A struct{}

func (a A) walk() {}

type B struct{}

func (b B) walk() {}

func main() {
	var i I
	i = A{}
	fmt.Printf("%T\n", i.(A))
	fmt.Printf("%T\n", i.(B)) // panic: interface conversion: main.I is main.A, not main.B
}
