package main

import "fmt"

var value1 float64

func main() {
	s := "abcd你"

	fmt.Println(s[4:] + "好")

	// 由于编译器行尾自动补全分号的缘故，加号必须放在第一行。
	str := "你好，" +
		"Go语言。"

	fmt.Println(str)
}

/*
返回结果：
你好
你好，Go语言。
*/
