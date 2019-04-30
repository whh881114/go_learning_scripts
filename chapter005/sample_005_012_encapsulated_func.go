package main

import (
	"fmt"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// nextNumber为一个函数，函数i为0
	nextNumber := getSequence()

	// 调用nextNumber函数，i变量自增1并返回
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())

	// 创建新的函数nextNumber1，并查看结果
	nextNumber1 := getSequence()
	fmt.Printf("%d ", nextNumber1())
	fmt.Printf("%d ", nextNumber1())
}

// 闭包会继承环境特性，在这个例子中，就是会保存i变量的值。
