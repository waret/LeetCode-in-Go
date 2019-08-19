package main

import "fmt"

func letterCombinations2(digits string) []string {


	if digits == "" {
		return nil
	}
	var m = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	sum := 1
	for i := 0; i < len(digits); i++ {
		sum *= len(m[digits[i]])
	}

	res := make([]string, sum)
	for i := 0; i < len(res); i++ {
		t := make([]byte, len(digits))
		x := 1
		for j := len(digits) - 1; j >= 0; j-- {
			for k := 0; k < len(m[digits[j]]); k++ {
				t[j] = m[digits[j]][(i/x)%(len(m[digits[j]]))]
			}
			x *= len(m[digits[j]])
		}
		res[i] = string(t)
	}
	return res
}

func letterCombinations(digits string) []string {
	var m = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	if digits == "" {
		return nil
	}

	ret := []string{""}

	// 让digits中所有的数字都有机会进行迭代。
	for i := 0; i < len(digits); i++ {
		var temp []string
		// 让ret中的每个元素都能添加新的字母。
		for j := 0; j < len(ret); j++ {
			// 把digits[i]所对应的字母，多次添加到ret[j]的末尾
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}
		ret = temp
	}

	return ret
}

func main() {
	fmt.Printf("%+v", letterCombinations2("23"))
}
