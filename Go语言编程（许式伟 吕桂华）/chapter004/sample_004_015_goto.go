package main

import "fmt"

func main() {
	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}

BREAK:
	fmt.Println("break")
}
