package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {

	tmp := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		tmp[i] = make([]byte, len(s))
	}

	x, y := 0, 0
	down := true
	for _, ch := range s {
		tmp[x][y] = byte(ch)
		if down && x < numRows-1 {
			x++
		} else if down && x == numRows-1 && x > 0 {
			down = false
			x--
			y++
		} else if !down && x > 0 {
			x--
			y++
		} else if !down && x == 0 {
			down = true
			x++
		} else {
			y++
		}
	}

	res := make([]byte, len(s))
	index := 0
	for i := 0; i < numRows; i++ {
		for _, ch := range tmp[i] {
			if ch > 0 {
				res[index] = ch
				index++
			}
		}
	}

	return string(res)
}

func convert2(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]bytes.Buffer, min(numRows, len(s)))

	curRow := 0
	goingDown := false
	for _, ch := range s {
		rows[curRow].WriteByte(byte(ch))
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}
	res := bytes.Buffer{}
	for _, r := range rows {
		res.Write(r.Bytes())
	}
	return res.String()
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func convert3(s string, numRows int) string {

	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace 步距
	p := numRows*2 - 2

	// 处理第一行
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// 处理中间的行
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// 处理最后一行
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}

func convert4(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ret := bytes.Buffer{}
	n := len(s)
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret.WriteByte(s[j+i])
			if i!=0 && i!=numRows-1 && j+cycleLen-i<n {
				ret.WriteByte(s[j+cycleLen-i])
			}
		}
	}
	return ret.String()
}

func main() {
	fmt.Println(convert4("AB", 1))
	fmt.Println(convert4("ABCDEFGHIJKLMNOPQRSTUVWX", 5))
}
