package main

/*
Go语言中的操作系统线程和goroutine的关系：
1- 一个操作系统线程对应用户态多个goroutine。
2- go程序可以同时使用多个操作系统线程。
3- goroutine和os线程是多对多的关系，即m:n。
*/

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
