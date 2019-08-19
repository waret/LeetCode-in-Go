package main

func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配 */

	dp[0][0] = true
	/**
	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
	 * 所以，需要把相应的 dp 设置成 true
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if p[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '.' || p[j] == s[i] {
				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				/* 此时，p[j] 的匹配情况与 p[j-1] 的内容相关。 */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					 * p[j] 无法与 s[i] 匹配上
					 * p[j-1:j+1] 只能被当做 ""
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					 * p[j] 与 s[i] 匹配上
					 * p[j-1;j+1] 作为 "x*", 可以有三种解释
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" 解释为 "" */
						dp[i+1][j] || /* "x*" 解释为 "x" */
						dp[i][j+1] /* "x*" 解释为 "xx..." */
				}
			}
		}
	}

	return dp[sSize][pSize]
}

func isMatch2(s, p string) bool {
	memo := make([][]Result, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		memo[i] = make([]Result, len(p)+1)
	}
	return dp(0, 0, s, p, memo)
}

type Result int

const (
	EMPTY Result = iota
	FALSE
	TRUE
)

func dp(i, j int, text, pattern string, memo [][]Result) bool {
	if memo[i][j] != EMPTY {
		return memo[i][j] == TRUE
	}
	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		firstMatch := i < len(text) && (pattern[j] == text[i] || pattern[j] == '.')
		if j+1 < len(pattern) && pattern[j+1] == '*' {
			ans = dp(i, j+2, text, pattern, memo) || firstMatch && dp(i+1, j, text, pattern, memo)
		} else {
			ans = firstMatch && dp(i+1, j+1, text, pattern, memo)
		}
	}
	if ans {
		memo[i][j] = TRUE
	} else {
		memo[i][j] = FALSE
	}
	return ans
}

func isMatch3(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func isMatch4(s, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		// pattern的第二个字符为*时，p[0]p[1]无需匹配任何s直接丢弃
		// 如果s[0]已经匹配成功的话，继续看pattern是否匹配s[1:]
		return isMatch4(s, p[2:]) || (firstMatch && isMatch4(s[1:], p))
	} else {
		return firstMatch && isMatch4(s[1:], p[1:])
	}
}

func (r Result) String() string {
	switch r {
	case EMPTY:
		return "EMPTY"
	case FALSE:
		return "FALSE"
	case TRUE:
		return "TRUE"
	default:
		return "UNKNOWN"
	}
}
