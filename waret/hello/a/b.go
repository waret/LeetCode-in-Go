package main

import "fmt"

type X1 interface {
	Error1() string
}

func New1(text string) X1 {
	// return object
	return xString{s: text}
}

func (x xString) Error1() string {
	return x.s
}

type X2 interface {
	Error2() string
}

func New2(text string) X2 {
	// return pointer
	return &xString{s: text}
}

func (x *xString) Error2() string {
	return x.s
}

type xString struct {
	s string
}

func main() {
	fmt.Println(New1("hello").Error1())
	fmt.Println(New1("hello"))
	fmt.Println(New2("world").Error2())
	fmt.Println(New2("world"))
}
