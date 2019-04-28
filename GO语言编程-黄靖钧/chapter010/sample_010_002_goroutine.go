package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") // 开一个新的goroutine执行
	say("hello")    // 当前goroutine执行
}

// 这个程序还不能正常运行，main()执行完后，就退出了，go执行里的协程还没有做完。
