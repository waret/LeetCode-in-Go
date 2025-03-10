package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		{
			p: para{
				one: []int{3, 2, 4},
				two: 8,
			},
			a: ans{
				one: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum0(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, twoSum1(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, twoSum2(p.one, p.two), "输入:%v", p)
		ast.Equal(a.one, twoSum3(p.one, p.two), "输入:%v", p)
	}
}
