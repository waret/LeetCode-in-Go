package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func readLeetCode() *leetcode {
	data, err := ioutil.ReadFile(leetCodeJSON)
	check(err)

	lc := new(leetcode)
	err = json.Unmarshal(data, lc)
	check(err)

	return lc
}

func (lc *leetcode) save() {

	if err := os.Remove(leetCodeJSON); err != nil {
		log.Panicf("删除 %s 失败，原因是：%s", leetCodeJSON, err)
	}

	raw, err := json.MarshalIndent(lc, "", "\t")
	if err != nil {
		log.Fatal("无法把Leetcode数据转换成[]bytes: ", err)
	}
	if err = ioutil.WriteFile(leetCodeJSON, raw, 0666); err != nil {
		log.Fatal("无法把 Marshal 后的 lc 保存到文件: ", err)
	}
	log.Println("最新的 LeetCode 记录已经保存。")
	return
}

func (lc *leetcode) ProgressTable() string {
	return lc.Record.progressTable()
}

func (lc *leetcode) AvailableTable() string {
	return lc.Problems.available().table()
}

func (lc *leetcode) FavoriteTable() string {
	return lc.Problems.favorite().table()
}

func (lc *leetcode) FavoriteCount() int {
	return len(lc.Problems.favorite())
}

func (lc *leetcode) UnavailableList() string {
	res := lc.Problems.unavailable().list()
	// 为了 README.md 文档的美观，需要删除最后一个换行符号
	return res[:len(res)-1]
}
