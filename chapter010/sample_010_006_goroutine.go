package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "纽约"
	ch <- "华盛顿"
	ch <- "伦敦"
	ch <- "北京"
	ch <- "东京"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s  ", input)
	}
}
