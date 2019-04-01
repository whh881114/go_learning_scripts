package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf("交换之前a的值为：%d\n", a)
	fmt.Printf("交换之前b的值为：%d\n", b)

	// 把a和b的地址传递给函数，然后a和b的地址交换，最终实现两个值的交换。
	// 这里只是一个示例，现在在go语言中，可以用a, b = b, a实现数字交换了。
	swap(&a, &b)

	fmt.Printf("交换之后a的值为：%d\n", a)
	fmt.Printf("交换之后b的值为：%d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
