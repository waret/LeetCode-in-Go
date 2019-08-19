package main

import (
	"fmt"
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
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0023(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{
				[]int{1, 4, 7},
				[]int{2, 5, 8},
				[]int{3, 6, 9},
			}},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},

		question{
			para{[][]int{
				[]int{1, 4, 7},
				[]int{2, 5, 8},
				[]int{},
			}},
			ans{[]int{1, 2, 4, 5, 7, 8}},
		},

		question{
			para{[][]int{
				[]int{1, 4, 7},
				[]int{},
				[]int{2, 5, 8},
			}},
			ans{[]int{1, 2, 4, 5, 7, 8}},
		},

		question{
			para{[][]int{}},
			ans{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(mergeKLists(ss2ls(p.one))), "输入:%v", p)
	}
}