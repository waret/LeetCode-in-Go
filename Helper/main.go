package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 程序辅助设置
const (
	VERSION = "8.0.0"
	USAGE   = `使用方法：
	helper readme
		获取最新的答题记录后，重新生成项目的 README.md 文件
	helper n
		生成第 n 题的答题文件夹
`
)

func main() {
	log.Printf("Hi, %s. I'm %s helper.\n", getConfig().Username, VERSION)

	if len(os.Args) != 2 {
		printUsageAndExit()
	}

	switch os.Args[1] {
	case "readme":
		buildReadme()
	default:
		problemNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			printUsageAndExit()
		}
		buildProblemFolder(problemNumber)
	}

}

func printUsageAndExit() {
	fmt.Println(USAGE)
	os.Exit(1)
}
