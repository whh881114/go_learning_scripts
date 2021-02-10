package main

import "fmt"

func hello() {
	fmt.Println("hello world!")
}

func main() {
	go hello()
	fmt.Println("hello main func.")
}
