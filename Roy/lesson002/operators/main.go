package main

import "fmt"

func main() {
	n1 := 19
	n2 := 3

	fmt.Println("==========================")
	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	n1--
	n2++
	fmt.Println(n1)
	fmt.Println(n2)

	fmt.Println("==========================")
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)

	fmt.Println("==========================")
	fmt.Println(n1 != n2 && n1 >= n2)
	fmt.Println(n1 != n2 || n1 < n2)

	fmt.Println("==========================")
	fmt.Printf("%b\n", n1)
	fmt.Printf("%b\n", n2)
	fmt.Printf("%b\n", n1|n2)
	fmt.Printf("%b\n", n1&n2)
}
