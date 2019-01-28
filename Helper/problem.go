package main

import (
	"fmt"
	"strings"
)

// TODO: 没有必要弄两个 problem 数据结构，可以合并
type problem struct {
	ID                                 int
	Title                              string
	TitleSlug                          string
	PassRate                           string
	Difficulty                         string
	IsAccepted, IsPaid, IsFavor, IsNew bool
	HasNoGoOption                      bool // 不能够使用 Go 语言解答
}

func newProblem(ps problemStatus) problem {
	level := []string{"", "Easy", "Medium", "Hard"}

	p := problem{
		ID:         ps.State.ID,
		Title:      ps.State.Title,
		TitleSlug:  ps.State.TitleSlug,
		PassRate:   fmt.Sprintf("%d%%", ps.ACs*100/ps.Submitted),
		Difficulty: level[ps.Difficulty.Level],
		IsAccepted: ps.Status == "ac",
		IsPaid:     ps.IsPaid,
		IsFavor:    ps.IsFavor,
		IsNew:      ps.State.IsNew,
	}

	return p
}

func (p problem) isAvailable() bool {
	if p.ID == 0 || p.IsPaid || p.HasNoGoOption {
		return false
	}
	return true
}

func (p problem) path() string {
	dir := "Algorithms"
	return fmt.Sprintf("./%s/%04d.%s", dir, p.ID, p.TitleSlug)
}

func (p problem) link() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
}

func (p problem) tableLine() string {
	// 题号
	res := fmt.Sprintf("|[%d](%s)|", p.ID, p.link())
	// 标题
	t := ""
	if p.IsAccepted {
		t = fmt.Sprintf(`[%s](%s)`, strings.TrimSpace(p.Title), p.path())
	} else {
		t = fmt.Sprintf(` * %s`, p.Title)
	}
	if p.IsNew {
		t += " :new: "
	}
	res += t + "|"
	// 通过率
	res += fmt.Sprintf("%s|", p.PassRate)
	// 难度
	res += fmt.Sprintf("%s|", p.Difficulty)
	// 收藏
	f := ""
	if p.IsFavor {
		f = "[❤](https://leetcode.com/list/oussv5j)"
	}
	res += fmt.Sprintf("%s|\n", f)
	return res
}

func (p problem) listLine() string {
	return fmt.Sprintf("- [%d.%s](%s)\n", p.ID, p.Title, p.link())
}
