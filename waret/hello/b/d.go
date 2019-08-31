package main

import "fmt"

// 1 2
func swap(a, b *int) {
	fmt.Printf("%p, %p\n", a, b)
	b, a = a, b
	fmt.Printf("%p, %p\n", a, b)
}

// 2 1
func swap2(a, b *int) {
	fmt.Printf("%p, %p\n", a, b)
	*b, *a = *a, *b
	fmt.Printf("%p, %p\n", a, b)
}

func main() {
	x, y := 1, 2
	fmt.Printf("%p, %p\n", &x, &y)
	swap2(&x, &y)
	fmt.Println(x, y)
	fmt.Printf("%p, %p\n", &x, &y)
}
