package main

import (
	"fmt"
	"time"
)

func main() {

	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)

	// 让玩家移动到3,1点
	p.MoveTo(Vec2{6, 0})

	// 如果没有到达就一直循环
	for !p.IsArrived() {
		select {
		case <-time.After(time.Second):
			// 打印每次移动后的玩家位置
			fmt.Println(p.Pos())
		}
	}

	p.MoveTo(Vec2{X: 6, Y: 8})

	// 如果没有到达就一直循环
	for !p.IsArrived() {
		select {
		case <-time.After(time.Second):
			// 打印每次移动后的玩家位置
			fmt.Println(p.Pos())
		}
	}

	p.Stop()

	select {
	case <-time.After(time.Second * 2):
	}

	fmt.Println("over!")
}
