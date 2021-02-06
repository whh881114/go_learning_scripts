package main

import "fmt"

// 自定义类型
// NewInt是一个新类型。
type NewInt int

// 类型别名：只存在代码编写过程中，代码编译之后根本不存在MyInt，仅提高代码的可读性。
type MyInit = int

func main() {
	var a NewInt
	a = 5
	fmt.Printf("Type: %T, Value: %d\n", a, a)

	var b MyInit
	b = 19
	fmt.Printf("Type: %T, Value: %d\n", b, b)
}
