package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
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
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: para{
				one: "advdfahjk",
			},
			a: ans{
				one: 7,
			},
		},
		{
			p: para{
				one: "dvdf",
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		t.Run(p.one, func(t *testing.T) {
			ast.Equal(a.one, lengthOfLongestSubstring4(p.one), "输入:%v", p)
		})
	}
}
