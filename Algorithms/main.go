package main

import (
	"fmt"
	"github.com/waret/LeetCode-in-Go/kit"
	"reflect"
)

func main() {
	a, b := 2, 3
	var c int
	c = kit.If(a > b, a, b).(int)
	fmt.Println(c)

	d, e := "aaa", "bbb"
	f := kit.If(d > e, d, e).(string)
	fmt.Println(reflect.TypeOf(f).String())
	fmt.Println(f)
}