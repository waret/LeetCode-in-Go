package main

import (
	"fmt"
)

func IsValidIpv4_1(src string) bool {

	t := 0
	for _, x := range src {
		if x == '.' {
			t++
		}
	}
	if t != 3 {
		return false
	}

	chars := []rune(src)

	start := 0
	end := 0
	for index, char := range src {
		if char == '.' || int(index) == len(chars) {
			sub := chars[start:end]
			sublen := len(sub)
			if sublen > 3 {
				return false
			}
			if sublen == 3 && sub[0] != '0' && sub[0] != '1' && sub[0] != '2' {
				return false
			}
			if sub[0] > '9' && sub[0] < '0' {
				return false
			}
			if sublen > 1 && sub[1] > '9' && sub[1] < '0' {
				return false
			}
			end++
			start = end
		} else {
			end++
		}
	}
	return true
}

func IsValidIpv4_2(src string) bool {
	chars := []rune(src)

	t := 0
	for _, x := range chars {
		if x == '.' {
			t++
		}
	}
	if t != 3 {
		return false
	}

	start := 0
	end := 0
	for index, char := range chars {
		if char == '.' || int(index) == len(chars) {
			sub := chars[start:end]

			s := 0
			for _, c := range sub {
				if c != ' ' && (c > '9' || c < '0') {
					return false
				}
				if c != ' ' {
					s++
				}
			}

			if s > 3 || s <= 0 {
				return false
			}

			ex := false
			for _, c := range sub {
				if c != ' ' {
					if s == 3 && ex == false {
						ex = true
						if c > '2' {
							return false
						}
					} else {
						if c > '9' && c < '0' {
							return false
						}
					}
				}
			}

			end++
			start = end
		} else {
			end++
		}
	}

	return true
}

func main() {
	//fmt.Println(IsValidIpv4_1("1.1.1.1"))
	//fmt.Println(IsValidIpv4_1("999.999.999.999"))
	fmt.Println(IsValidIpv4_2("1.1.1.1"))
	fmt.Println(IsValidIpv4_2("999.999.999.999"))
	fmt.Println(IsValidIpv4_2("  9 9. 99.  9 9 . 9 9   "))
}
