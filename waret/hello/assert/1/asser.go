package main

import "fmt"

type I interface {
	walk()
	quack()
}

type S struct{}

func (s S) walk()  {}
func (s S) quack() {}

func main() {
	var i I
	i = S{}
	val, ok := i.(interface {
		walk()
	})
	fmt.Println(val, ok)
}
