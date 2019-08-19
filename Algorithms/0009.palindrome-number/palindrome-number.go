package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	for x >= y {
		y = 10*y + x%10
		x = x / 10
		if x > 0 && y == 0 {
			return false
		}
		if x == 0 || x == y || x/10 == y {
			return true
		}
	}
	return false
}

func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	if x < 10 {
		return true
	}

	y := 0
	for x >= y {
		y = 10*y + x%10
		x = x / 10
		if x == y || x/10 == y {
			return true
		}
	}
	return false
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	isPalindrome(0)
}
