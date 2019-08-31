package t

import "fmt"

type X struct {
	Name string
}

type Y struct {
	X
	Age int32
}

type SayHi interface {
	Hello()
}

func (x *X) Hello() {
	fmt.Println("x.Name: ", x.Name)
}

func (y *Y) Hello() {
	y.X.Hello()
	fmt.Println("y.Age: ", y.Age)
}
