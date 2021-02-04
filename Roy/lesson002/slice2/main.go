package main

import "fmt"

func main() {
	var a []string
	var b []int
	var c = []bool{false, true}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("=======================================")

	// 基于数组得到切片
	a1 := [5]int{55, 56, 57, 58, 59}
	b1 := a1[1:4]
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)
	fmt.Printf("数组a1的地址：%p\n", &a1) // a1是变量，需要使用&来取变量地址。
	fmt.Printf("数组b1的地址：%p\n", b1)
	fmt.Println("=======================================")

	// a切片再次切片
	c1 := b1[0:len(b1)]
	fmt.Printf("%T\n", c1)
	fmt.Println("=======================================")

	// make函数构造切片
	d1 := make([]int, 5, 10)
	fmt.Println(d1, len(d1), cap(d1))
}
