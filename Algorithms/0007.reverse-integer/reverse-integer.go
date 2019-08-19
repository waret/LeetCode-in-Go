package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	y := 0
	for x != 0 {
		pop := x % 10
		if (y > math.MaxInt32/10) || (y == math.MaxInt32/10 && pop > 7) || (y < math.MinInt32/10) || (y == math.MinInt32/10 && pop < -8) {
			return 0
		}
		y = 10*y + pop
		x /= 10
	}
	return y
}

// 限于64位机器
func reverse2(x int) int {
	y := 0
	if x == 0 {
		return x
	}

	for x != 0 {
		y = 10*y + x%10
		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return y
}

func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func reverse3(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	// 还原 x 的符号到 res
	res = sign * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func main() {
	fmt.Println(reverse(-321))
}
