package main

import "fmt"

var value1 float64

func main() {
	var v1 complex64
	v1 = 3.2 + 12i
	v2 := 3.2 + 12i        // v2隐式声明，默认是complex128类型
	v3 := complex(3.2, 12) // v3数值等于v2
	v := v2 + v3
	fmt.Println(v1, v2, v3, v)

	vr := real(v)
	vi := imag(v)

	fmt.Println(vr, vi)
}
