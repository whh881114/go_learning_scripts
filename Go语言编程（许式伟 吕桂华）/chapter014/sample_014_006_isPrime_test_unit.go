package main

import "fmt"

func main() {
	fmt.Println(isPrime(25))
}

// 大于1的自然数中，除了1和它本身以外不再有其他因数的数称为质数
func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}

	if value%2 == 0 || value%3 == 0 {
		return false
	}

	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {

			return false
		}
	}
	return true
}
