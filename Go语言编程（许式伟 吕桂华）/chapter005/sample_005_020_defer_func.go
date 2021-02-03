package main

import "fmt"

func main() {
	b()
}

func a() {
	defer un(trace("a")) // 初始化defer函数的参数，所以输出trace()的结果
	fmt.Println("a的逻辑代码")
}

func b() {
	defer un(trace("b")) // 初始化defer函数的参数，所以输出trace()的结果
	fmt.Println("b的逻辑代码")
	a()
}

func trace(s string) string {
	fmt.Println("开始执行", s)
	return s
}

func un(s string) {
	fmt.Println("结束执行")
}

/*
1. defer函数先要初始化，所以在执行b()函数后，然后再执行a()函数，所在返回结果先是：
*******************
开始执行 b
b的逻辑代码
*******************
开始执行 a
a的逻辑代码
*******************
结束执行		// 函数b()执行的defer结果
*******************
结束执行		// 函数a()执行的defer结果
*/
