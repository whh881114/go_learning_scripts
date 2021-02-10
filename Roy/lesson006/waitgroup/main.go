package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello world!")
	wg.Done()
}

func main() {
	go hello()
	wg.Add(1) // 1. 创建一个goroutine；2. 在新的goroutine中执行hello函数。
	fmt.Println("Hello main func.")
	wg.Wait() // 阻塞，一直等待所有的goroutine函数执行完。
}
