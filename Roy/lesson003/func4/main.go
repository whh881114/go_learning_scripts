package main

import "fmt"

// panic和recover

func a() {
	fmt.Println("func a")
}

// 1. recover()必须搭配defer使用。
// 2. defer一定要在可能引发panic的语句之前定义。
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	panic("panic in b") // 程序异常退出前要调用下defer语句。
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
