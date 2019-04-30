package main

import "fmt"

func main() {
	a := 100
	if a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a大于20\n")
	}
	fmt.Printf("a的值是：%d\n", a)

	fmt.Println("*******************************")
	a = 11
	if a < 20 {
		fmt.Printf("a小于20\n")

		if a > 10 {
			fmt.Println("a大于10")
		} else {
			fmt.Println("a小于10")
		}
	} else {
		fmt.Printf("a大于20\n")
	}
	fmt.Printf("a的值是：%d\n", a)

	fmt.Println("*******************************")
	a = 11
	if a > 20 {
		fmt.Printf("a大于20\n")
	} else if a < 10 {
		fmt.Printf("a小于10\n")
	} else {
		fmt.Printf("a大于10\n")
		fmt.Printf("a小于20\n")
	}
	fmt.Printf("a的值是：%d\n", a)

	fmt.Println("*******************************")
	a = 13
	if a > 20 {
		fmt.Printf("a大于20\n")
	} else if a < 10 {
		fmt.Printf("a小于10\n")
	} else if a == 11 {
		fmt.Printf("a等于11\n")
	} else {
		fmt.Printf("a大于10\n")
		fmt.Printf("a大于20\n")
		fmt.Printf("a不等于11\n")
	}
	fmt.Printf("a的值是：%d\n", a)
}
