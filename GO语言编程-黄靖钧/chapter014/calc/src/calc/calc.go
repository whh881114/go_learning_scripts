package main

import (
	"fmt"
	"goTraining/GO语言编程-黄靖钧"
	"os"
	"strconv"
)

// 将一个匿名函数赋值给一个变量
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println("\tadd\tAddition of two values.")
	fmt.Println("\tsqrt\tSquare root of a non-negative float value.")
}

func main() {
	args := os.Args

	// Debug:
	// for i, v := range args {
	// 	fmt.Printf("Index: %d    ----> Value: %s\n", i, v)
	// }

	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1> <integer2>")
			return
		}

		ret := GO语言编程_黄靖钧.Add(v1, v2)
		fmt.Println("Result: ", ret)

	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := GO语言编程_黄靖钧.Sqrt(v)
		fmt.Println("Result: ", ret)

	default:
		Usage()
	}
}
