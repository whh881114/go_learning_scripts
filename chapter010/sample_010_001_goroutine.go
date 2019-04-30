package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("这里是main()开始的地方。")
	// 使用go关键字来并发运行两个协程。
	go longWait()
	go shortWait()

	fmt.Println("挂起main()")
	// 挂起工作时间以纳秒(ns)为单位
	time.Sleep(10 * 1e9)

	fmt.Println("这里是main()结束的地方。")
}

func longWait() {
	fmt.Println("开始longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("结束longWait()")
}

func shortWait() {
	fmt.Println("shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("shortWait()")
}
