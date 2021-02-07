package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/tmp/whh.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := "hello world!"
	file.WriteString(str)
	defer file.Close()
}
