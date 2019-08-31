package main

import "fmt"

func main() {
	b := "helliiiiiiiio"
	fmt.Println(b)
	ptr := &b
	fmt.Println(b[0])
	fmt.Printf("%p, %p, %p\n", ptr, &ptr, &b)
	*ptr = "world"
	fmt.Printf("%p, %p, %p\n", ptr, &ptr, &b)

	c := int64(10)
	cptr := &c
	fmt.Printf("%p, %p, %p\n", &c, &cptr, cptr)
	*cptr = 30
	fmt.Printf("%p, %p, %p\n", &c, &cptr, cptr)

	fmt.Println()
	d := []int32{1, 2, 3}
	dptr := &d
	// 0xc0000aa000, 0xc0000aa004, 0xc0000aa008, 0xc0000a8000, 0xc0000ac000
	fmt.Printf("%p, %p, %p, %p, %p", &d[0], &d[1], &d[2], dptr, &dptr)
}
