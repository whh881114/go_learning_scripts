package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)

	b := &a
	fmt.Printf("Type: %T, Value: %v。\n", b, b)

	c := "豪杰"
	fmt.Printf("&c=%v\n", &c)
	// b = &c // 取c的内存地址不能赋值给b，指针是带类型的，所以字符串的地址不能赋值给整型的地址。

	d := 100
	b = &d
	fmt.Println(b)

	// *取地址对应的值
	fmt.Println(*b)
}
