package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("/tmp/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
}
