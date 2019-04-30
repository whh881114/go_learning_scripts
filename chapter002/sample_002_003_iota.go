package main

import "fmt"

const (
	a       = iota             // 0
	b                          // 1
	c                          // 2
	d, e, f = iota, iota, iota // d = 3, e = 3, f = 3，同一行值相同
	g       = iota             // g = 4，显示赋值
	h       = "h"              // 单独赋值，但是iota依旧递增，此时为5
	i                          // i = "h"，默认使用上面的赋值，iota依旧递增，此时为6。
	j       = iota             // j = 7
)

const z = iota // z = 0，每个单独定的const常量中，iota都会重置，此时z=0。

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, z)
}
