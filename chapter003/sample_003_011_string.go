package main

import "fmt"

func main() {
	s := "Hello 世界！"
	b := []byte(s) // 转换为[]byte，自动复制数据
	b[5] = ','     // 修改[]byte

	// b[5]的赋值时要用单引号，此时值的类型为byte，如果是双引号表示string。

	fmt.Printf("%s\n", s) // s不能被修改，内容保持不变
	fmt.Printf("%s\n", b) // 修改后的数据
}
