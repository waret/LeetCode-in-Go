package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aQuaYi/GoKit"
)

const (
	unavailableFile = "unavailable.json"
	leetCodeJSON    = "leetcode.json"
)

type leetcode struct {
	Username string    // 用户名
	Ranking  int       // 网站排名
	Updated  time.Time // 数据更新时间

	Record   record   // 已解答题目与全部题目的数量，按照难度统计
	Problems problems // 所有问题的集合
}

func updateMyData() {
	lc := latestLeetCode()
	lc.save()
}

func latestLeetCode() *leetcode {
	problems, record := parseAlgorithms(getAlgorithms())
	lc := &leetcode{
		Username: getConfig().Username,

		Record:   record,
		Problems: *problems,

		Ranking: getRanking(),
		Updated: time.Now(),
	}

	return lc
}

func parseAlgorithms(alg *algorithms) (*problems, record) {
	hasNoGoOption := readUnavailable()
	myProblems := &problems{}
	r := record{}

	for _, ps := range alg.Problems {
		p := newProblem(ps)
		if hasNoGoOption[p.ID] {
			p.HasNoGoOption = true
		}
		myProblems.add(p)
		r.update(p)
	}

	return myProblems, r
}

func readUnavailable() map[int]bool {
	type unavailable struct {
		List []int
	}

	if !GoKit.Exist(unavailableFile) {
		log.Panicf("%s 不存在，没有不能解答的题目", unavailableFile)
	}

	raw := read(unavailableFile)
	u := unavailable{}
	if err := json.Unmarshal(raw, &u); err != nil {
		log.Panicf("获取 %s 失败：%s", unavailableFile, err)
	}

	res := make(map[int]bool, len(u.List))
	for i := range u.List {
		res[u.List[i]] = true
	}

	return res
}
func getAlgorithms() *algorithms {
	URL := "https://leetcode.com/api/problems/Algorithms/"

	raw := getRaw(URL)

	res := new(algorithms)
	if err := json.Unmarshal(raw, res); err != nil {
		log.Panicf("无法把 json 转换成 Category: %s\n", err.Error())
	}

	// 如果，没有登录的话，也能获取数据，但是用户名，就不是本人
	if res.User != getConfig().Username {
		log.Fatal("没有获取到本人的数据")
	}

	return res
}

// getRanking 让这个方法优雅一点
func getRanking() int {
	// 获取网页数据
	URL := fmt.Sprintf("https://leetcode.com/%s/", getConfig().Username)
	data := getRaw(URL)
	str := string(data)

	// 通过不断裁剪 str 获取排名信息

	i := strings.Index(str, "ng-init")
	j := i + strings.Index(str[i:], "ng-cloak")
	str = str[i:j]

	i = strings.Index(str, "(")
	j = strings.Index(str, ")")
	str = str[i:j]

	//	fmt.Println("2\n", str)

	strs := strings.Split(str, ",")
	str = strs[6]

	//	fmt.Println("1\n", str)

	i = strings.Index(str, "'")
	j = 2 + strings.Index(str[2:], "'")

	//	fmt.Println("0\n", str)

	str = str[i+1 : j]

	r, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("无法把 %s 转换成数字Ranking", str)
	}

	return r
}
