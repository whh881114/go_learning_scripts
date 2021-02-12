package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex     // 互斥锁
var rwLock sync.RWMutex // 读写互斥锁：多个goroutine同时讯加的是读锁，写的时候加的是写锁。

func read() {
	defer wg.Done()
	// 互斥锁
	lock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()

	// // 读写互斥锁
	// rwLock.RLock()
	// fmt.Println(x)
	// time.Sleep(time.Millisecond * 10)
	// rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	// 互斥锁
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 50)
	lock.Unlock()

	// // 读写互斥锁
	// rwLock.Lock()
	// x = x + 1
	// time.Sleep(time.Millisecond * 50)
	// rwLock.Unlock()
}

func main() {
	start := time.Now()
	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Printf("\n运行时间：%v\n", end.Sub(start))
}
