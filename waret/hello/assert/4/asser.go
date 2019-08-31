package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)
import "strconv"

type Stringer interface {
	String() string
}

type Number struct {
	v int
}

func (number *Number) String() string {
	return strconv.Itoa(number.v)
}

func main() {
	n := &Number{1}
	switch v := interface{}(n).(type) {
	case Stringer:
		fmt.Println("Stringer:", v)
	default:
		fmt.Println("Unknown")
	}

	var i, j interface{}
	i = Number{2}

	// i 为实例，但实例未实现 Type 接口（实例指针实现的）
	v1, o1 := i.(Stringer)
	fmt.Println("v1, o1", v1, o1) // false

	// i 为实例，但 Type 为指针
	v12, o12 := i.(*Stringer)
	fmt.Println("v12, o12", v12, o12) // false

	// i 为实例，Type 为实例
	v2, o2 := i.(Number)
	fmt.Println("v2, o2", v2, o2) // true

	// i 为实例，Type 为指针
	v21, o21 := i.(*Number)
	fmt.Println("v21, o21", v21, o21) // false

	j = &Number{2}

	// j 为实例指针，实例指针实现了 Type 接口
	v13, o13 := j.(Stringer)
	fmt.Println("v13, o13", v13, o13) // true

	var s Stringer
	s = &Number{3}

	//var x *Stringer
	//x = &s  // 可以赋值，但无法断言
	//v31, o31 := x.(*Stringer)
	//fmt.Println("v31, o31", v31, o31)

	v3, o3 := s.(Stringer)
	fmt.Println("v3, o3", v3, o3) // true

	v4, o4 := s.(*Number)
	fmt.Println("v4, o4", v4, o4) // true

	//v41, o41 := s.(Number) // 无法赋值，s 一定为指针类型
	//fmt.Println("v41, o41", v41, o41) // true

	var w io.Writer
	w = os.Stdout
	v5, o5 := w.(io.ReadWriter)
	fmt.Println("v5, o5", v5, o5) // true

	w = new(bytes.Buffer)
	v6, o6 := w.(io.ReadWriter)
	fmt.Println("v6, o6", v6, o6) // true
}
