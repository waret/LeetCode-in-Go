package main

import (
	"fmt"
	"reflect"
)

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 当hlen等于nlen的时候，需要i == 0
	for i := 0; i <= hlen-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}

func strStr1(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i < hlen-nlen; i++ {
		fmt.Println(reflect.TypeOf(haystack[i : i+nlen]))
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr1("hello", "ll"))
	fmt.Println("hello"[1:3] == "hello"[1:3])
}
