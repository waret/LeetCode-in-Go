package main

import "fmt"

type A struct {
	ax, ay int
	string
}
type B struct {
	A
	x A
	string
	ax, ay int
}

func main() {
	b := B{A{1, 2, "bb"}, A{4, 5, "xx"}, "aa", 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.x, b.A, b.A.ax, b.A.ay, b.string, b.A.string)
	fmt.Println(b.A)
}
