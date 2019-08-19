package main

import (
	"fmt"
	"strings"
	"unicode"
    "log"
)

var where = log.Print

func main() {
    where()
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
    where()
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
    where()

    a := [...]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("Array item", i, "is", a[i])
    }
}
