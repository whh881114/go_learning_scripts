package main

import "fmt"

func main() {
	a := 20
	ap := &a

	fmt.Printf("a的地址：%X\n", &a)
	fmt.Printf("ap所指向的地址：%X\n", ap)
	fmt.Printf("ap所指向的地址所存的变量值：%d\n", *ap)
	fmt.Printf("ap指针变量在内存中申请的地址为：%X\n", &ap)

	fmt.Println("*********************************************")
	a += 1
	fmt.Printf("a的值：%d\n", a)
	fmt.Printf("ap的值：%d\n", *ap)
}
