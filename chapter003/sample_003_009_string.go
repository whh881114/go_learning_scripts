package main

import "fmt"

var value1 float64

func main() {
	s := "你"
	t := "好"

	// 字符串比较大小的判断是根据第一字节来判定的。
	if s < t {
		fmt.Println(s[0], t[0])
		fmt.Println(s[1], t[1])
		fmt.Println(s[2], t[2])
	}

	a := "a"
	b := "b"

	if a < b {
		fmt.Println(a[0], "小于", b[0])
	}
}
