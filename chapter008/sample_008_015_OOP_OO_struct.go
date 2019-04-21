package main

import "fmt"

type Integer int

// 面向对象，调用方法是foo.bar()格式。
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 面向过程，调用方法是bar(foo)格式。
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1

	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	if Integer_Less(a, 2) {
		fmt.Println(a, "Less 2")
	}
}
