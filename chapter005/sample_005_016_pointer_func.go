package main

import "fmt"

func main() {
	n := 0
	res := &n
	multiply(2, 4, res)
	fmt.Println("Result: ", *res)
}

func multiply(a, b int, res *int) {
	*res = a * b
}
