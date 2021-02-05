package main

import "fmt"

func sayHello() {
	fmt.Println("Hello 沙河小王子。")
}

func sayHello2(name string) {
	fmt.Println("Hello", name)
}

// 定义接收多个参数函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

func intSum2(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func intSum3(a ...int) int {
	ret := 0
	for _, i := range a {
		ret += i
	}
	return ret
}

func intSum4(a int, b ...int) int {
	ret := a
	for _, i := range b {
		ret += i
	}
	return ret
}

// Go语言中函数参数类型简写
func intSum5(a, b int) int {
	ret := a + b
	return ret
}

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sayHello()
	sayHello2("汪浩浩")
	fmt.Println(intSum(1, 2))
	fmt.Println(intSum2(1, 3))
	fmt.Println(intSum3(1, 3, 4, 5, 8, 100, 9))
	fmt.Println(intSum4(100, 3, 4, 5, 8, 100, 9))
	fmt.Println(calc(100, 3))
}
