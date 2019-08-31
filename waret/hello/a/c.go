package main

import (
	"fmt"
	"reflect"
)

type C interface {
	hello() string
}

type cString struct {
	s string
}

func (c *cString) hello() string {
	return c.s
}

func New(text string) C {
	return &cString{s: text}
}

func xxx(c *C) {
	fmt.Printf("1: %v\n", (*c).hello())
	*c = New("hello")
	fmt.Printf("2: %v\n", (*c).hello())
	*c = &cString{s: "xxx"}
	fmt.Printf("5: %v\n", (*c).hello())
}

func main() {
	c := New("xxxxxxx")
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("3: %v\n", c.hello())
	xxx(&c)
	fmt.Printf("4: %v\n", c.hello())
}
