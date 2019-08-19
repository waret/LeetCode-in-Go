package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func printListNode(ls *ListNode) {
	if ls.Next == nil {
		fmt.Print(ls.Val)
		return
	}
	printListNode(ls.Next)
	fmt.Print(ls.Val)
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{7, 0, 8}),
			},
		},
		{
			p: para{
				one: makeListNode([]int{9, 8, 7, 6, 5}),
				two: makeListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans{
				one: makeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		{
			p: para{
				one: makeListNode([]int{0}),
				two: makeListNode([]int{5, 6, 4}),
			},
			a: ans{
				one: makeListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		aa := addTwoNumbers2(p.one, p.two)
		printListNode(p.one)
		fmt.Print(" + ")
		printListNode(p.two)
		fmt.Print(" = ")
		printListNode(aa)
		fmt.Println()
		ast.Equal(a.one, aa, "输入:%v", p)
	}
}

// go test -bench=. -benchmem -run=none
func Benchmark_addTwoNumbers(b *testing.B) {

	funcs := []struct {
		name string
		f    func(l1 *ListNode, l2 *ListNode) *ListNode
	}{
		{"addTwoNumbers1", addTwoNumbers1},
		{"addTwoNumbers2", addTwoNumbers2},
	}

	for _, f := range funcs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f.f(makeListNode([]int{2, 4, 3}), makeListNode([]int{5, 6, 4}))
			}
		})
	}
}
