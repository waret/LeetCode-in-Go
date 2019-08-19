package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	p para
	a ans
}

type para struct {
	one int
}

type ans struct {
	one bool
}

func Test_Problem0009(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 12321,
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: 1231,
			},
			a: ans{
				one: false,
			},
		},
		{
			p: para{
				one: -12321,
			},
			a: ans{
				one: false,
			},
		},
		{
			p: para{
				one: 0,
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: 1,
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		t.Run(fmt.Sprintf("%d", p.one), func(t *testing.T) {
			ast.Equal(a.one, isPalindrome(p.one), "输入:%v", p)
		})
	}
}
