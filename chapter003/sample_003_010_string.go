package main

import "fmt"

var value1 float64

func main() {
	s := "我是中国人"
	fmt.Printf("字符串s的长度是：%d。\n", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("<%d %c>\n", i, s[i])
		//fmt.Printf("%d = %v\n", i, s[i]) // 输出单字节值
	}

	fmt.Printf("\n")

	for _, v := range s {
		fmt.Printf("%c", v)
	}
	fmt.Print("\n")
}

// 中文每个字符可能由多个字切组成，遍历列表长度时打印出的值会是乱码。
