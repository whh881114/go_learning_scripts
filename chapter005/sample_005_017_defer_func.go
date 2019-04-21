package main

import "fmt"

func main() {
	fmt.Println("return: ", a()) // 打印结果为return: 0
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defter2: ", i) // 打印结果为defter2: 2
	}()

	defer func() {
		i++
		fmt.Println("defter1: ", i) // 打印结果为defter1: 1
	}()

	return i
}

// 当main()函数中的fmt.Println要执行时，就会运行a()这个函数，然后执行里面的return语句前的函数，然后按defter逆序执行。
