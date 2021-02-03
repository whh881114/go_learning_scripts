package main

import (
	"fmt"
	"time"
)

func main() {
	// genErr()
	throwsPanic(genErr)
}

func genErr() {
	fmt.Println(time.Now(), "正常的语句")

	defer func() {
		fmt.Println(time.Now(), "defer 正常语句")
		panic("第二个错误")
	}()
	panic("第一个错误")
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(time.Now(), "捕获到异常：", r)
			b = true
		}
	}()

	f()
	return
}

// 看不懂，之后再学习吧。
