package main

import "fmt"

const pi = 3.14 // 声明常量格式。

// 指量声明常量
const (
	a = 100
	b = 1000 // c和d的值与b一样。
	c
	d
)

// iota 枚举
/*
	iota是go语言的常量计数器，只能在常量的表达式中使用。
	iota在const关键字出现时将被重置为0，const中每新增一行常量将使用iota计数一次。
	iota可理解为const语句块中的行索引。
*/
const (
	aa = iota // 0
	bb        // 1	const声明如果不写，默认就和上一行一样。bb = iota
	cc        // 2
	dd        // 3
)

const (
	n1 = iota // 0
	n2        // 1
	_         // 2
	n4        // 3
)

const (
	nn1 = iota // 0
	nn2 = 100
	nn3 = iota // 2
	nn4        // 3
)

const nn5 = iota // 0

const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<10 <=> 2^10
	MB = 1 << (10 * iota) // 1<<20 <=> 2^20
	GB = 1 << (10 * iota) // 1<<30 <=> 2^30
	TB = 1 << (10 * iota) // 1<<40 <=> 2^40
	PB = 1 << (10 * iota) // 1<<50 <=> 2^50
)

const (
	a2, b2 = iota + 1, iota + 2 // iota=0, a2=1, b2=2
	c2, d2                      // iota=1, c2=2, d2=3
	e2, f2                      // iota=2, e2=3, d2=4
)

func main() {
	// pi = 888.8 // 常量声明后，其值不可以改变。
	fmt.Println(pi)
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(n1, n2, n4)
	fmt.Println(nn1, nn2, nn3, nn4, nn5)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a2, b2, c2, d2, e2, f2)
}
