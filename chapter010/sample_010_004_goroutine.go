package main

import "fmt"

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
	// 原始代码的print内容是在ch后面
	// fmt.Println("Counting")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

/*
第一个for就是和main()并发运行Count()函数，先往channel里写值，并打印内容。
此时是并发状态运行，10个一起运行，由于是先写channel里的值，再打印内容，就存在一种情况，channel已经写完了，
与此同时main()函数中的第二个for得知channel处于非阻塞状态就可以读了，读完了main()函数就退出了，而在Count()里的print可能没打印完。
*/
