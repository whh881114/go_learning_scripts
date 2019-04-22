package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("/tmp/goDir", 0777)
	os.MkdirAll("/tmp/goDir2/1/2/3/4", 0755)

	err := os.Remove("/tmp/goDir")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("/tmp/goDir2")
}
