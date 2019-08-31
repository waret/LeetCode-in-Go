package main

import (
	"fmt"
	"math"
)

const Pi64 float64 = math.Pi

var x float32 = float32(Pi64)
var y float64 = Pi64
var z complex128 = complex128(Pi64)
var m complex128 = complex(Pi64, Pi64)

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

type XPlayer = Player

type mint int32
type ptr *mint
type pptr *ptr

func main() {
	xx := new(mint)
	yy := ptr(xx)
	zz := pptr(&yy)
	//mm := pptr(&ptr(xx))
	nn := mint(10)
	//oo := &mint(10)
	fmt.Printf("%T, %T, %T, %T\n", xx, yy, zz, nn)

	ins := XPlayer{Name: "aa", HealthPoint: 1, MagicPoint: 2}
	fmt.Println(ins)
	fmt.Printf("%T\n", ins)
	fmt.Println(math.Pi)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%v, %v, %v\n", real(m), imag(m), m)
}
