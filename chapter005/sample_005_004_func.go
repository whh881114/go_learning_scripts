package main

import "fmt"

// pipe输入值为一个匿名函数，匿名函数返回值是一个整型，然后呢，这个返回值作为输入，返回一个int。
// 听着有点拗口，简单地说，pipe函数的输入值是要经过一个式子算出来的结果传输入，只是恰巧这个式子叫匿名函数而已。
func pipe(ff func() int) int {
	return ff() // 返回值刚好是个匿名函数算出来的结果，可以这么理解，传进去的是什么值就返回什么值。
}

// formatFunc定义函数类型
// 使用type来声明FormatFunc为一个函数， 输入是一个string和两个int，返回一个string。
type FormatFunc func(s string, x, y int) string

// 定义format函数，输入参数为一个FormatFunc函数，一个string和两个int，返回一个string。
func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}

func main() {
	// 直接匿名函数当参数，返回值是100，打印出的结果也是100。
	s1 := pipe(func() int { return 100 })

	// 第一个参数为FormatFunc的匿名函数：func(s string, x, y int) string { return fmt.Sprintf(s, x, y) }
	// 第二个参数是："%d, %d"
	// 第三和第四参数是： x, y
	// 计算过程是就是，第二，三和四这三个参数传入第一个匿名函数，然后返回这三个参数，打印出来的结果是：10, 20。
	s2 := format(func(s string, x, y int) string { return fmt.Sprintf(s, x, y) }, "%d, %d", 10, 20)
	fmt.Println(s1, s2)
}
