package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		println(k, v)
	}

	fmt.Println("******************************")
	numbers := []int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("第%d次，x的值为%d。\n", i, x)
	}

	fmt.Println("******************************")
	numbers2 := [5]int{1, 2, 3, 4}
	for i, x := range numbers2 {
		fmt.Printf("第%d次，x的值为%d。\n", i, x)
	}
}
