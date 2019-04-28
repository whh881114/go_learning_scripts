package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c > 10 {
			break
		}
	}
}

// 采用共享内存的方式来判断协程是否完成。此例中的就是c这个变量，也就是全局变量counter。
