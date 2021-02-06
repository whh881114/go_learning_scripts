package main

import "fmt"

// 可以给任意类型追加方法，不能给go语言中已有的类型添加方法。

type MyInt int

func (m *MyInt) sayHi() {
	fmt.Println("Hello MyInt ~")
}

func main() {
	var a MyInt
	a.sayHi()
}
