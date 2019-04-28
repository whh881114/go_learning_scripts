package main

import "fmt"

func main() {
	my_anon_func := func(x, y int) int {
		return x + y
	}

	fmt.Printf("匿名函数直直接赋值给一个变量：%d。\n", my_anon_func(1, 2))
	fmt.Println()

	func() {
		fmt.Println("匿名函数使用示例，自己声明，然后自已调用。")
	}()
}
