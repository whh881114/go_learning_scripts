package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello world!")
	wg.Done() // 计数器-1
}

func main() {
	go hello() // 1. 创建一个goroutine；2. 在新的goroutine中执行hello函数。
	wg.Add(1)  // 计数器+1，根据协程的数量设置。
	fmt.Println("Hello main func.")
	wg.Wait() // 阻塞，一直等待所有的goroutine函数执行完。
}
