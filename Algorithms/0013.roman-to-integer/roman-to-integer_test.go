package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{"XXXIX"},
			ans{39},
		},
		{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		{
			para{"MCMLXXVI"},
			ans{1976},
		},
		{
			para{"MMMCMXCIX"},
			ans{3999},
		},
		{
			para{"LXIX"},
			ans{69},
		},
		{
			para{"XLIX"},
			ans{49},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}
}
