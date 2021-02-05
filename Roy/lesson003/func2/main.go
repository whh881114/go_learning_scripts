package main

import "fmt"

// 定义全局变量
var num = 10

// 定义函数
func testGlobal() {
	fmt.Println("全局变量", num) // 可以在函数中访问全局变量

	num := 100
	fmt.Println("局部变量", num)
}

func add(x, y int) int {
	return x + y
}

// calc函数，接收三个参数，其中前两个为整型，后一个为函数，这个函数是接收两个整型，返回一个整型。
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	testGlobal()

	// i变量此时只能在for循环的语句中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		在 JavaScript 中总是见到一等公民字样，一直很纳闷，什么是一等公民？？？

		用必应词典翻译一下，英文叫first-class citizen，再维基百科一下

		In programming language design, a first-class citizen (also type, object, entity, or value) in a given
		programming language is an entity which supports all the operations generally available to other entities. These operations
		typically include being passed as an argument, returned from a function, and assigned to a variable.
	*/

	// 将函数赋值给变量
	abc := testGlobal
	abc()

	fmt.Println(calc(100, 200, add))
}
