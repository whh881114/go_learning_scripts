package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// runtime.GOMAXPROCS(1) // 设置我的Go程序只有一个逻辑核心。当启用时，结果是先打印完a()再打印完b()，或者是先b()再a()。
	// 当不使用时，就是跑在所有核心时，其结果就是a()和b()随机打印。
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
