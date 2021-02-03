package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff

	fmt.Println(a, b, c)              // 默认10进制输出
	fmt.Printf("%b %b %b\n", a, b, c) // 2进制输出
	fmt.Printf("%O %O %O\n", a, b, c) // 8进制输出
	fmt.Printf("%X %X %X\n", a, b, c) // 16进制输出

	fmt.Printf("%p\n", &c) // 变量内存地址

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex64
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 字符串转义
	fmt.Println("\"c:\\go\"")

	var s1 = "单行文本"
	// ``中间的文件会原样输出。
	var s2 = `这
	是
	多行
	文本。`

	fmt.Println(s1)
	fmt.Println(s2)
}
