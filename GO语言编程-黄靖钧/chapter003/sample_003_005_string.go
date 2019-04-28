package main

import "fmt"

var value1 float64

func main() {
	str1 := `苟利国家生死以\n
	岂因祸福避趋之`

	str2 := "今天天气\n真好"

	fmt.Println(str1)
	fmt.Println(str2)
}

// 反引号，反引号可以跨行，并且引号内的所有内容都会直接输出，包括转义字符和空格缩进等。
