package main

import "fmt"

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int 64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	MyPrintf(v1, v2, v3, v4)
}

// 这里重点说下空接口，也就是这个interface{}，接口是定义methods方法，空接口是就没有没有任何method，所以任何变type都实现了这个功能。
// 当前MyPrintf()函数接收类型为...interface{}，所以任何值都能传进去，就相当于是一个大容量可以存很多东西。
// 空接口是后续的内容。
