package main

import (
	"fmt"
	"strings"
)

// 定义一个函数，它的返回值是一个函数。
func a() func() {
	name := "沙河娜扎"
	return func() {
		// fmt.Println("hello, 沙河小王子。")
		fmt.Println("hello", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= 1
		return base
	}
	return add, sub
}

func main() {
	// 匿名函数就是没有函数名的函数，匿名函数多用于实现回调函数和闭包。
	// func (参数) (返回值) {
	// 	函数体
	// }
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
	r := a() // a()返回一个匿名函数传给r，r此时是一个闭包。
	r()

	x := makeSuffixFunc(".gz")
	fmt.Println(x("httpd-2.4.2.tar.gz"))

	x = makeSuffixFunc(".xxx")
	fmt.Println(x("httpd-2.4.2.tar.gz"))

	y, z := calc(10)
	fmt.Println(y(10))
	fmt.Println(z(10))
}
