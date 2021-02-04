package main

import "fmt"

func main() {
	var a = 100
	var b = "沙河娜扎"
	var c = false

	fmt.Println(a, b, c)
	fmt.Printf("a=%v\n", a) // %v俗称占位符
	fmt.Printf("a=%T\n", a)
}
