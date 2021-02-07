package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "Hello world!\n"
	err := ioutil.WriteFile("/tmp/whh_ioutil.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
