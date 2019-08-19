package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one []string
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
				one: "babad",
			},
			a: ans{
				one: []string{"bab", "aba"},
			},
		},
		{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: []string{"bb"},
			},
		},
		{
			p: para{
				one: "abbcccddcccbba",
			},
			a: ans{
				one: []string{"abbcccddcccbba"},
			},
		},
		{
			p: para{
				one: "a",
			},
			a: ans{
				one: []string{"a"},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Contains(a.one, longestPalindrome3(p.one), "输入:%v", p)
	}
}