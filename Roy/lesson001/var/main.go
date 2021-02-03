package _var

import "fmt"

// 全局变量，必须要用var进行声明。
var alex = "dsb"

// 定义两个返回值的函数。
func foo() (string, int) {
	return "alex", 9000
}

func main() {
	fmt.Println("========================")
	fmt.Printf("全局变量：%s\n", alex)

	// 声明的变量必须要使用到。
	var name string
	var age int
	fmt.Println("========================")
	fmt.Println(name)
	fmt.Println(age)

	var (
		a string // ""
		b int    // 0
		c bool   // false
		d string // ""
	)
	a = "沙河"
	b = 100
	c = true
	d = "100"
	fmt.Println("========================")
	fmt.Println(a, b, c, d)

	var x string = "老男孩"
	fmt.Println("========================")
	fmt.Println(x)
	fmt.Printf("%s嘿嘿嘿\n", x)

	var y = 200 // 编译器会根据变量初始值判断变量类型。
	var z = true
	fmt.Println("========================")
	fmt.Println(y, z)

	// 简短声明，只能在函数内部使用。
	nazha := "哇哈哈"
	fmt.Println("========================")
	fmt.Println(nazha)

	// 调用foo函数，使用一个占位符，"_"，用于接收不需要的变量值。
	aa, _ := foo()
	fmt.Println("========================")
	fmt.Println(aa)
}
