package main

import (
	"fmt"
	"math"
	"time"
)

// select多路复用

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f1:%d", i)
		time.Sleep(time.Millisecond * 50)
	}
}

func f2(ch chan string) {
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f2:%d", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go f1(ch1)
	go f2(ch2)

	// 只要某个channel中有值即可以取到。
	for {
		select {
		case ret := <-ch1:
			fmt.Println(ret)
		case ret := <-ch2:
			fmt.Println(ret)
		default:
			fmt.Println("暂时取不到值。")
			time.Sleep(time.Second)
		}
	}
}
