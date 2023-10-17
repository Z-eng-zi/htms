package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a int = 0
	maxNum := 100
	rand.Seed(time.Now().UnixNano()) //  设置随机数种子
	secretNumber := rand.Intn(maxNum)
	b := secretNumber
	fmt.Println("现在进入猜数字游戏")
	if a != b {
		fmt.Print("请输入你猜测的整数：")
		fmt.Scanln(&a)
		fmt.Println("你猜测的数字：", a)
		if a == b {
			fmt.Println("恭喜你，猜对了，游戏结束！")
		} else {
			fmt.Println("不好意思，你猜错了!")
			fmt.Println("正确答案是 ", secretNumber)
		}

	}
}