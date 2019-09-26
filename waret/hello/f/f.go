package main

import (
	"fmt"
	"net/url"
	"unsafe"
)

type Address struct {
	Address string
	Zip     float64
}

type User struct {
	Name    string
	Age     int32
	Ahel    bool
	Test    float64
	Address Address
}

//type Error struct {
//	errCode uint8
//}
//
//func checkError(err error) {
//	fmt.Println(&err == nil)
//	fmt.Println(err == nil)
//	fmt.Println(err)
//	fmt.Println(&err)
//	if err != nil {
//		fmt.Println("hello")
//		panic(err)
//	}
//}
//
//func (e *Error) Error() string {
//	switch e.errCode {
//	case 1:
//		return "file not found"
//	case 2:
//		return "time out"
//	case 3:
//		return "permission denied"
//	default:
//		return "unknown error"
//	}
//}

func main() {
	var a = [5]User{{Name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, {Name: "bbbbbbbbbbbbbbbbbb"}, {Name: "cccccccccccccccccccccc"}}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%p, %p, %p, %p, %p, %v, %v, %v\n", &a, &a[i], &a[i].Name, &a[i].Test, &a[i].Address.Zip, a[i].Name, a[i].Age, a[i].Address)
		fmt.Printf("%v\n", unsafe.Sizeof(a[i].Age))
	}
	b := a[1:3]
	fmt.Println(b)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%p, %p, %p, %p, %p, %v, %v, %v\n", &b, &b[i], &b[i].Name, &b[i].Test, &b[i].Address.Zip, b[i].Name, b[i].Age, b[i].Address)
	}
	c4 := append(b, User{Name: "dddddddddd", Age: 10})
	fmt.Println(c4)
	for i := 0; i < len(c4); i++ {
		fmt.Printf("%p, %p, %p, %p, %p, %v, %v, %v\n", &c4, &c4[i], &c4[i].Name, &c4[i].Test, &c4[i].Address.Zip, c4[i].Name, c4[i].Age, c4[i].Address)
	}
	fmt.Println(b)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%p, %p, %p, %p, %p, %v, %v, %v\n", &b, &b[i], &b[i].Name, &b[i].Test, &b[i].Address.Zip, b[i].Name, b[i].Age, b[i].Address)
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%p, %p, %p, %p, %p, %v, %v, %v\n", &a, &a[i], &a[i].Name, &a[i].Test, &a[i].Address.Zip, a[i].Name, a[i].Age, a[i].Address)
	}

	var numbers []string
	for i := 0; i < 10; i++ {
		numbers = append(numbers, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%p, %v\n", &numbers[i], numbers[i])
	}

	var boolean bool
	fmt.Printf("boolean: %v, &boolean: %p, boolean==false: %t\n", boolean, &boolean, boolean == false)

	var integer int
	fmt.Printf("integer: %v, &integer: %p, integer==0: %t\n", integer, &integer, integer == 0)

	var str string
	fmt.Printf("str: %v, &str: %p, str==\"\": %t\n", str, &str, str == "")

	var array [3]int
	fmt.Printf("array: %v, &array: %p, array==[3]int{}: %t\n", array, &array, array == [3]int{})

	var s []int
	fmt.Printf("s: %v, &s: %p, s==nil: %t\n", s, &s, s == nil)
	var s2 = []int{}
	fmt.Printf("s2: %v, &s2: %p, s2==nil: %t\n", s2, &s2, s2 == nil)

	var m map[int]string
	fmt.Printf("m: %v, &m: %p, m==nil: %t\n", m, &m, m == nil)
	var m2 = map[int]string{}
	fmt.Printf("m2: %v, &m2: %p, m2==nil: %t\n", m2, &m2, m2 == nil)

	var c chan int
	fmt.Printf("c: %v, &c: %p, c==nil: %t\n", c, &c, c == nil)
	var c2 = make(chan int)
	fmt.Printf("c2: %v, &c2: %p, c2==nil: %t\n", c2, &c2, c2 == nil)

	var e url.Error
	fmt.Printf("e: %v, &e: %p, e==nil: %t\n", e, &e, e == url.Error{})
	var ep *url.Error
	fmt.Printf("ep: %v, &ep: %p, ep==Error{}: %t\n", ep, &ep, ep == nil)

	var x error
	fmt.Printf("x: %v, &x: %p, x==nil: %t\n", x, &x, x == nil)
	x = ep
	fmt.Printf("x: %v, &x: %p, x==nil: %t\n", x, &x, x == nil)

	var x2 interface{}
	fmt.Printf("x2: %v, &x2: %p, x2==nil: %t\n", x2, &x2, x2 == nil)
	x2 = e
	fmt.Printf("x2: %v, &x2: %p, x2==nil: %t\n", x2, &x2, x2 == nil)

}
