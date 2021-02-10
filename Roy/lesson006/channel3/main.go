package main

import "fmt"

// 接收值时判断通道是否关闭。

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	// 利用for循环去通道ch1中接收值
	for {
		ret, ok := <-ch1 // 当通道关闭的时候ok=false
		if !ok {
			break
		}
		fmt.Println(ret)
	}
	fmt.Println("===========================================")

	var ch2 = make(chan int, 100)
	go send(ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
