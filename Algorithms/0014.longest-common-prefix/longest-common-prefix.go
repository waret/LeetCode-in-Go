package main

func longestCommonPrefix(strs []string) string {
	s := shortest(strs)
	for i := range s {
		for _, ss := range strs {
			if ss[i] != s[i] {
				if i == 0 {
					return ""
				} else {
					return string(s[:i])
				}
			}
		}
	}
	return s
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s := strs[0]
	for _, c := range strs {
		if len(c) < len(s) {
			s = c
		}
	}
	return s
}

func longestCommonPrefix2(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func main() {
	longestCommonPrefix([]string{"abcdd", "abcde"})
}
