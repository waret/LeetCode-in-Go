package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	sign, abs := clean(s)
	x := 0
	for _, b := range abs {
		x = 10*x + int(b-'0')
		if x > 1<<31-1 {
			if sign > 0 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}
	}

	return sign * x
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	switch {
	case s[0] >= '0' && s[0] <= '9':
		sign, abs = 1, s
	case s[0] == '-':
		sign, abs = -1, s[1:]
	case s[0] == '+':
		sign, abs = 1, s[1:]
	default:
		return
	}
	for i, b := range abs {
		if b > '9' || b < '0' {
			abs = abs[:i]
			break
		}
	}
	return
}


func myAtoi2(s string) int {
	result := int64(0)
	negative := false
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) {
		if s[i] == '-' {
			negative = true
			i++
		} else if s[i] == '+' {
			negative = false
			i++
		}
	}
	for i < len(s) && s[i] >= '0' && s[i] <='9' && result <= math.MaxInt32 {
		result *= 10
		result += int64(s[i] - '0')
		i++
	}
	if negative {
		result = -result
	}
	if result >= math.MaxInt32 {
		return math.MaxInt32
	} else if result <= math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

func main() {
	//fmt.Println(myAtoi2("123"))
	//fmt.Println(myAtoi2("-2147483649"))
	//fmt.Println(myAtoi2("-2147483647"))
	//fmt.Println(myAtoi2("1"))
	fmt.Println(myAtoi2("42"))
}
