package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "纽约"
	ch <- "华盛顿"
	ch <- "伦敦"
	ch <- "北京"
	ch <- "东京"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s  ", input)
	}
}
