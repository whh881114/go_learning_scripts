package main

import "fmt"

func main() {
	a := 20
	c := 200

	c = a // c = 20
	fmt.Println("赋值操作，把a赋值给c，所以c的值为：", c)

	c += a // c = 40
	fmt.Println("相加和赋值运算符，实际为c = c + a，所以c的值为：", c)

	c -= a // c = 20
	fmt.Println("相减和赋值运算符，实际为c = c -a，所以c的值为：", c)

	c /= a // c = 1
	fmt.Println("相除和赋值运算符，实际为c = c / a，所以c的值为：", c)

	c <<= 2 // c = 4
	fmt.Println("左移和赋值运算符，所以c的值为：", c)

	c >>= 2 // c = 1
	fmt.Println("右移和赋值运算符，所以c的值为：", c)

	c &= 2 // c = 0
	fmt.Println("按位与和赋值运算符，所以c的值为：", c)

	c ^= 2 // c = 2
	fmt.Println("按位异或和赋值运算符，所以c的值为：", c)

	c |= 2 // c = 2
	fmt.Println("按位或和赋值运算符，所以c的值为：", c)
}
