package main

import "fmt"

var value1 float64

func main() {
	s := "aA你2"
	fmt.Println("字符串长度：", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// 编译错误，字符串的值不能修改。
	s = "你好，"
	t := s

	s += "世界。" // 字符串可以连接，但原子符串不会改变。
	fmt.Println(s)

	// s[0] = "L"
	fmt.Println(t) // 原字符串没有改变。
}

/*
返回结果：
字符串长度： 6
97	// a
65	// A
228	// 你
189	// 你
160	// 你
50	// 2
你好，世界。
你好，
*/
