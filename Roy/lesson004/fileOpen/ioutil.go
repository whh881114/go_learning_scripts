package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("/Users/roy/bin/auto_scaling_groups.sh")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(content))
}
