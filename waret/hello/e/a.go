package main

import (
	"fmt"
	"os"
	"reflect"
)

type I interface {
	SayHi()
}

type S struct {
	name string
}

func (s *S) SayHi() {
	fmt.Println(s.name)
}

func (s S) String() string {
	return fmt.Sprintf("hello %s", s.name)
}

//func (s S) SayHi() {
//	fmt.Println(s.name)
//}

func main() {
	s := S{name: "world"}
	fmt.Printf("%p, %v\n", &s, s)
	s = S{name: "waret"}
	fmt.Printf("%p, %v\n", &s, s)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	sp := &S{name: "world"}
	fmt.Printf("%p, %p, %v\n", &sp, sp, sp)
	sp = &S{name: "waret"}
	fmt.Printf("%p, %p, %v\n", &sp, sp, sp)
	fmt.Println(sp)
	fmt.Println(reflect.TypeOf(sp))

	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err))   // "true"
}
