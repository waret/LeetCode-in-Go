package main

import (
	"fmt"
	"github.com/waret/go-study/waret/hello/d/t"
)

func main() {
	x := t.X{Name: "waret"}
	y := t.Y{
		X: t.X{
			Name: "li",
		},
		Age: 10,
	}

	var w interface{}
	w = y

	switch w.(type) {
	case t.X:
		fmt.Println("t.X")
	case t.Y:
		fmt.Println("t.Y")
		fmt.Println(w.(t.Y))
	case t.SayHi:
		fmt.Println("t.SayHi")
	default:
		fmt.Println("default")
	}

	v, ok := w.(t.SayHi)
	fmt.Println(v, ok)
	x.Hello()
	y.Hello()
}
