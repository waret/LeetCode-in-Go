package main

import "fmt"

func mySqrt2(x int) int {
	res := x
	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	s, e := 0, x/2
	m := (s + e) / 2
	for true {
		if m*m <= x {
			if (m+1)*(m+1) > x {
				return m
			} else {
				s = m + 1
			}
		} else {
			e = m - 1
		}
		m = (s + e) / 2
	}
	return 1
}

func mySqrt3(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x/2
	for l <= r {
		m := l + (r-l)/2
		if m*m < x {
			l = m + 1
		} else if m*m > x {
			r = m - 1
		} else {
			return m
		}
	}
	return r
}

func main() {
	fmt.Println(mySqrt(1))
}
