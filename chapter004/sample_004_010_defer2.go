package main

import "fmt"

var i = 0

func print() {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		defer print()
	}
}

// 返回的是5个5，这是因为每个deter都在函数轮询之后，最后才执行。
//5
//5
//5
//5
//5
