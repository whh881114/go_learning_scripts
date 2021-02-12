package main

import "fmt"

func main() {
	// 声明一个存放int类型，容量为1的通道
	var ch = make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case ch <- i: // 第一次可以赋值0，当i为1时，存不进去，因为容量满了，所以就执行下一步，即打印结果。
		case ret := <-ch:
			fmt.Println(ret)
		}
	}
}
