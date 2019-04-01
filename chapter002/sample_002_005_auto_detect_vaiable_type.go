package main

import (
	"fmt"
	"os"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {
	fmt.Printf("%q, %q, %q", HOME, USER, GOROOT)
}
