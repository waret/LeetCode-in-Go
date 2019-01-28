package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// read 负责读取文件
// 这是一个通用的方法
func read(path string) []byte {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	return data
}

func write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	check(err)
}

// 利用 VSCode 打开文件
func vscode(filename string) {
	open("code", "-r", filename)
}

func chrome(link string) {
	open("google-chrome", link)
}

func open(software string, arg ...string) {
	cmd := exec.Command(software, arg...)
	_, err := cmd.Output()
	check(err)
}

// TODO: 把所有的 if err 修改成 check
func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
