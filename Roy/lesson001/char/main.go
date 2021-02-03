package main

import (
	"fmt"
)

func main() {
	s1 := "Golang"
	c1 := 'G'
	fmt.Println(s1, c1)

	s2 := "中国"
	c2 := '中' // UTF-8编码下，一个中文占3个字节
	fmt.Println(s2, c2)

	s3 := "hello沙河"
	fmt.Println(len(s3))

	// 遍历字符串
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%d --> %c\n", i, s3[i])
	}

	fmt.Println("==============")

	// 使用range遍历，是按照rune类型去遍历。
	for k, v := range s3 {
		fmt.Printf("%d --> %c\n", k, v)
	}

	fmt.Println("==============")

	// 强制类型转换
	s5 := "big"
	byteArray := []byte(s5)
	fmt.Println(byteArray)
	byteArray[0] = 'p'
	fmt.Println(string(byteArray))

	// 作业题：hello --> olloh
	// 方法一
	s6 := "hello"
	s7 := []byte(s6)
	s8 := []byte(s6)
	for i := 0; i < len(s7); i++ {
		s8[i] = s7[len(s7)-i-1]
	}
	fmt.Println(string(s8))

	// 方法二
	s9 := "hello"
	s10 := ""
	byteArrayS9 := []byte(s9)
	for i := len(byteArrayS9) - 1; i >= 0; i-- {
		s10 = s10 + string(byteArrayS9[i])
	}
	fmt.Println(s10)

	// 方法三
	s11 := "hello"
	s12 := []byte(s11)
	for i := 0; i < len(s12)/2; i++ {
		s12[i], s12[len(s12)-1-i] = s12[len(s12)-1-i], s12[i]
	}
	fmt.Println(string(s12))
}
