package main

import "fmt"

func main() {
	if a := 10; a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a的值是：", a)
	}

	// a是在if代码块中定义的，所以不能在函数中调用。
	// fmt.Println("a的值是：", a)
}
