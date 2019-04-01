package main

import "fmt"

func main() {
	s := "Hello 世界！"
	r := []rune(s) // 转换为[]rune，自动复制数据

	fmt.Printf("s字符串长度为：%d。\n", len(s))
	fmt.Printf("r字符串长度为：%d。\n", len(r))

	r[6] = '中' // 修改[]rune
	r[7] = '国' // 修改[]rune

	fmt.Println(s)         // s不能被修改，内容保持不变
	fmt.Println(string(r)) // 转换为字符串，又一次复制数据
}

// 转换成rune类型后，就可以按字符个数来数位置了。
// 返回结果：
// s字符串长度为：15。
// r字符串长度为：9。
// Hello 世界！
// Hello 中国！
