package problem0344

import "fmt"

func reverseString1(s string) string {
	bytes := []byte(s)

	i, j := 0, len(s)-1

	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}

func reverseString2(s []byte)  {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseString3(s []byte) {
	for start := 0; start < len(s) / 2; start++ {
		s[start], s[len(s) - 1 - start] = s[len(s) - 1 - start], s[start]
	}
}

func printReverseString(s []byte) {
	printReverseStringHelper(0, s)
}

func printReverseStringHelper(index int, s []byte) {
	if s == nil || index >= len(s) {
		return
	}
	printReverseStringHelper(index + 1, s)
	fmt.Print(string(s[index]))
}

func main() {
	s := []byte{'h','e','l','l','o'}
	reverseString2(s)
	for _, v := range s {
		fmt.Print(string(v))
	}
	fmt.Println()
	reverseString2(s)
	for _, c := range s {
		fmt.Print(string(c))
	}
	fmt.Println()

	printReverseString(s)
}