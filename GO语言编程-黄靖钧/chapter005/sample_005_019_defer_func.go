package main

import (
	"fmt"
	"time"
)

func main() {
	i, p := a()
	fmt.Println("return: ", i, p, time.Now())
}

func a() (i int, p *int) {
	// i++
	// 下面有参数的defer函数，函数中的i与外围子函数的i并不是同一个参数，初始化时i值为0，打印结果为0。
	defer func(i int) {
		fmt.Println("defer3: ", i, &i, time.Now())
	}(i)

	defer fmt.Println("defer2: ", i, &i, time.Now())

	defer func() {
		fmt.Println("defer1: ", i, &i, time.Now())
	}()

	i++
	defer func() {
		fmt.Println("print1: ", i, &i, time.Now())
	}()

	fmt.Println("print2: ", i, &i, time.Now())
	return i, &i
}
