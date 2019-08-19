package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	// 所有符号都添加完毕
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// "(" 不用担心匹配问题，
	// 只要 left > 0 就可以直接添加
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	// 想要添加 ")" 时
	// 需要 left < right，
	// 即在 bytes[:idx] 至少有一个 "(" 可以与 这个 ")" 匹配
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

func generateParenthesis1(n int) []string {
	var ans []string
	backtrack(&ans, "", 0, 0, n)
	return ans
}

func backtrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(ans, cur+")", open, close+1, max)
	}
}

func generateParenthesis2(n int) []string {
	var ans []string
	if n == 0 {
		return append(ans, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generateParenthesis2(c) {
				for _, right := range generateParenthesis2(n - 1 - c) {
					ans = append(ans, "("+left+")"+right)
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(generateParenthesis2(3))
}
