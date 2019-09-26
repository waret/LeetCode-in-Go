package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s1 := "debug"
	s2 := "all:debug"

	fmt.Println(s1[strings.Index(s1, ":")+1:])
	fmt.Println(s2[strings.Index(s2, ":")+1:])
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s2))
}
