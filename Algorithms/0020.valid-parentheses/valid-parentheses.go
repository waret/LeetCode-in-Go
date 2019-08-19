package main

func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // '('+1 is ')'
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}

func isValid1(s string) bool {
	size := len(s)
	half := size / 2
	top := 0
	stack := make([]byte, half+1)
	for i := 0; i < size; i++ {
		switch s[i] {
		case '(':
			stack[top] = ')'
			top++
		case '{':
			stack[top] = '}'
			top++
		case '[':
			stack[top] = ']'
			top++
		case ')', '}', ']':
			if top > 0 && s[i] == stack[top-1] {
				top--
			} else {
				return false
			}
		}
		if top > half {
			return false
		}
	}
	return top == 0
}
