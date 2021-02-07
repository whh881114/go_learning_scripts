package main

import (
	"fmt"
	"time"
)

func tickDemo() {
	// 定义一个1秒间隔的定时器
	ticker := time.Tick(time.Second)

	// 每秒都会执行的任务
	for i := range ticker {
		fmt.Println(i)
	}
}

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%#v\n", now)
	fmt.Println("========================================")

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println("========================================")

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println("========================================")

	var timestamp int64
	timestamp = 1612664970
	timeObject := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObject)
	year := timeObject.Year()
	month := timeObject.Month()
	day := timeObject.Day()
	hour := timeObject.Hour()
	minute := timeObject.Minute()
	second := timeObject.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println("========================================")

	// 定时器
	// tickDemo()

	// 时间格式化，格式化的模板为Go的出生时间2006年1月2号15点04分。
	// 真是一个奇葩。
	// 2006-01-02 15:04:05 --> 便于记忆：2006 1 2 3 4 5
	// 2006 --> year
	// 01 --> month
	// 02 --> day
	// 15 --> hour
	// 04 --> minute
	// 05 --> second
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("15:04:05 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}
