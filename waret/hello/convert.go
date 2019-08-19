package main

import "fmt"

type Y interface {
	y()
}

type X struct {
	x int64
}

func main() {

	aa := new(X)
	aa.x = 1000

	bb := X{x: 2222}

	var m interface{}

	m = aa // 此时m为指针
	if obj, ok := m.(*X); ok {
		fmt.Printf("xxxxxxxxx: %d \n", obj.x)
	} else {
		fmt.Println(ok)
	}

	m = bb // 此时m为对象
	if obj, ok := m.(X); ok {
		fmt.Printf("xxxxxxxxx: %d \n", obj.x)
	} else {
		fmt.Println(ok)
	}

}
