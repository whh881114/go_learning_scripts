package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go语言中内置的map并发不是安全的。

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=%v, v=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
