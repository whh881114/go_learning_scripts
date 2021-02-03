package main

import "fmt"

func main() {
	age := 18
	for age > 0 {
		fmt.Println(age)
		age--
	}

	fmt.Println("==========")

	finger := 1
	switch finger {
	case 1:
		fmt.Println("You are the one.")
	case 2, 3, 4:
		fmt.Println("No...")
	case 5:
		fmt.Println("Again...")
	default:
		fmt.Println("Wahaha")
	}

	fmt.Println("==========")

	finger2 := 3
	switch {
	case finger2 == 1:
		fmt.Println("You are the one.")
	case finger2 >= 2 && finger2 <= 4:
		fmt.Println("No...")
		fallthrough
	case finger2 >= 5:
		fmt.Println("Again...")
	}
}
