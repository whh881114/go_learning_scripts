package main

import (
	"fmt"
)

func main() {
	x := 3
	y := 4
	xPLUSy, xTIMEy := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMEy)
}

// SumAndProduct 返回A+B和A*B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
