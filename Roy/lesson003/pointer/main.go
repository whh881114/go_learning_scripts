package main

/*
指针和地址有什么区别？
地址：就是内存地址，用字节来描述的内存地址。
指针：指针是带类型的。

&：表示取地址
*：根据地址取值
*/

import (
	"fmt"
)

func modifyArray(a1 [3]int) {
	a1[0] = 100
}

func modifyArray2(a1 *[3]int) {
	a1[0] = 100
}

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

	// 指针的应用
	a1 := [3]int{1, 2, 3}
	modifyArray(a1)
	fmt.Println(a1)

	modifyArray2(&a1)
	fmt.Println(a1)
}
