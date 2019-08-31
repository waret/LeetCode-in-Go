package main

import (
	"fmt"
	"reflect"
)

type Fragment int

func (f Fragment) Exec() {
	fmt.Printf("%p\n", &f)
	f = 10
}
func (f *Fragment) Exec2() {
	fmt.Printf("%p\n", f)
	*f = 10
}
func main() {
	fragment := Fragment(1)
	fmt.Println(reflect.TypeOf(fragment))
	fmt.Printf("%p --  %v \n", &fragment, fragment)
	fragment.Exec()
	fmt.Printf("%p --  %v \n", &fragment, fragment)
	(&fragment).Exec()
	fmt.Printf("%p --  %v \n", &fragment, fragment)
	fragment.Exec2()
	fmt.Printf("%p --  %v \n", &fragment, fragment)
	(&fragment).Exec2()
	fmt.Printf("%p --  %v \n", &fragment, fragment)

	fmt.Println("----------------------------")

	fragment2 := new(Fragment)
	*fragment2 = 1
	fmt.Println(reflect.TypeOf(fragment2))
	fmt.Printf("%p --  %v \n", fragment2, *fragment2)
	fragment2.Exec()
	fmt.Printf("%p --  %v \n", fragment2, *fragment2)
	(*fragment2).Exec()
	fmt.Printf("%p --  %v \n", fragment2, *fragment2)
	fragment2.Exec2()
	fmt.Printf("%p --  %v \n", fragment2, *fragment2)
	(*fragment2).Exec2()
	fmt.Printf("%p --  %v \n", fragment2, *fragment2)
}
