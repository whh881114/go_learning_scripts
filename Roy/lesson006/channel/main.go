package main

import "fmt"

func main() {
	var ch1 chan int    // 定义的一个channel，只能传递int类型的数据。
	var ch2 chan string // 定义的一个channel，只能传递string类型的数据。

	// channel是引用类型
	fmt.Println("ch1:", ch1)
	fmt.Println("ch2:", ch2)

	// make函数初始化（分配内存）：slice，map和channel。
	ch3 := make(chan int, 10)
	fmt.Println("ch3:", ch3)

	// channel的操作类型：发送，接收和关闭。
	// 发送和接收的符号：<-

	ch3 <- 10 // 把10发送到ch3中
	// <-ch3        // 从ch3中接收值，直接丢弃。
	ret := <-ch3 // 从ch3中接收值，保存了到变量ret中。
	fmt.Println(ret)

	// 关闭通道
	close(ch3)
	// 关闭channel后，可以取channel里面的值。

	// 1- 关闭的通首再接收，能取到对应类型的零值。
	// 2- 往关闭的通道中发送值会引发panic。
	// 3- 关闭一个已关闭的channel会引发panic。
	ret2 := <-ch3
	fmt.Println(ret2)
}
