package main

import "fmt"

var i = 0

func print(i int) {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		defer print(i)
	}

}

// 返回结果：defer是一个栈，遵循先入后出，所以结果是倒序。
//4
//3
//2
//1
//0
