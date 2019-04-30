package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time // Time是time包中一个struct结构体。
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}

	// 调用匿名Time上的String方法
	fmt.Println("完整的时间格式：", m.String())

	// 调用myTime.first3Chars
	fmt.Println("前三个字符：", m.first3Chars())
}
