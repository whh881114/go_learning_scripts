package main

import "fmt"

func main() {
	for a := 0; a < 5; a++ { // a在for中定义，其他地方是用不到的，所以下面需要再次声明a
		fmt.Printf("a的值是：%d\n", a)
	}

	fmt.Println("********************************")
	a := 0
	b := 5
	for a < b {
		fmt.Printf("a的值是：%d\n", a)
		a++
	}

	fmt.Println("********************************")
	str := "abcz"
	for i, char := range str {
		fmt.Printf("字符串第%d个字符的值为%d\n", i, char)
	}

	fmt.Println()
	for _, char := range str { //	忽略第二个值。
		fmt.Println(char)
	}

	fmt.Println()
	for range str { // 忽略全部返回值，只执行下面代码块。
		fmt.Println("执行成功！")
	}
}
