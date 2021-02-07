package main

import (
	"fmt"
	"time"
)

// time包练习题
// 编写一个函数，接收Time类型的参数，函数内部将传进来的时间格式化输出2017/06/19 20:30:05格式
// 编写程序统计一段代码的执行耗时时间，单位精确到微秒。

func main() {
	now := time.Now()
	fmt.Println("-- 获取当前时间，格式化输出2017/06/19 20:30:05格式。")
	fmt.Println("--", now.Format("2006/01/02 15:04:05"))

	start := time.Now().UnixNano() / 1000
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Println("程序运行时间（微秒）：", end-start)
}
