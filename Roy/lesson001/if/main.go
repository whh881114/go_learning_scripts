package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else if age < 18 {
		fmt.Println("Warning...")
	} else {
		fmt.Println("成年了！")
	}

	// 以下写法中，age2的作用域仅存在于if判断中。
	if age2 := 28; age2 > 18 {
		fmt.Println("成年了！")
	}
	// fmt.Println(age2)，age2已经不能再被引用了。
}
