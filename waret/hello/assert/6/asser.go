package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	A
	Age int
}

func main() {
	var i interface{}
	i = B{A: A{Name: "xx"}, Age: 1}
	v1, o1 := i.(A)
	fmt.Println(v1, o1)

	i = A{Name: "yy"}
	v2, o2 := i.(B)
	fmt.Println(v2, o2)
}
