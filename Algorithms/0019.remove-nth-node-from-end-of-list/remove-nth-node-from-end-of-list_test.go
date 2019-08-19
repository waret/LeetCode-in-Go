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
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0019(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2}, 1},
			ans{[]int{1}},
		},
		question{
			para{[]int{1, 2, 3, 4, 5}, 1},
			ans{[]int{1, 2, 3, 4}},
		},
		question{
			para{[]int{1, 2, 3, 4, 5}, 2},
			ans{[]int{1, 2, 3, 5}},
		},
		question{
			para{[]int{1}, 1},
			ans{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(removeNthFromEnd(s2l(p.one), p.two)), "输入:%v", p)
	}
}


// go test -bench=. -benchmem -run=none
func Benchmark_aaa(b *testing.B) {

	funcs := []struct {
		name string
		f    func(head *ListNode, n int) *ListNode
	}{
		{"removeNthFromEnd", removeNthFromEnd},
		{"removeNthFromEnd1", removeNthFromEnd1},
		{"removeNthFromEnd2", removeNthFromEnd2},
		{"removeNthFromEnd3", removeNthFromEnd3},
	}

	for _, f := range funcs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := make([]int, 999999)
				for i := 0; i < 999999; i++ {
					s[i] = i
				}
				l := s2l(s)
				f.f(l, 2)
			}
		})
	}
}