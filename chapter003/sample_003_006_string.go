package main

import "fmt"

var value1 float64

func main() {
	s := "abcd你"

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	//c := s[:len(s)]
	c := s[len(s)-1]
	fmt.Println(len(s), c)

	fmt.Println(s[0:7])
	// fmt.Println(s[4:]) // 返回：你
	// fmt.Println(s[:3]) // 返回：abc
}
