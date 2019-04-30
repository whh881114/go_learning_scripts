package main

import "fmt"

// 演示go语言中函数是使用的是值传递与址传递用法

type Integer int

func (a *Integer) Add(b Integer) {
	*a += b
}

func (a Integer) Add2(b Integer) {
	a += b
}

func main() {
	var a Integer = 1
	fmt.Println("### 调用add2方法（值传递）并不会影响main函数中的a的值，输出结果应该为1。")
	a.Add2(2)
	fmt.Printf("a = %d\n\n", a)

	fmt.Println("### 调用add方法（址传递）这会影响main函数中的a的值，输出结果应该为3。")
	a.Add(2)
	fmt.Printf("a = %d\n", a)
}
