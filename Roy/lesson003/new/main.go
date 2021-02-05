package main

import "fmt"

// new，用来初始化值类型指针的。
// make，用来初始化slice，map，chan。

func main() {
	var a = new(int)
	*a = 10
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int)
	c[0] = 3
	fmt.Println(*c)
}
