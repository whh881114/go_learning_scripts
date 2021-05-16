package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var pathSep string = string(os.PathSeparator)

func listAll(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			fmt.Println(path + pathSep + fi.Name())
		} else {
			fmt.Println(path + pathSep + fi.Name())
		}
	}
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	listAll(root)
}
