package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog 中文"
	strSli := strings.Fields(str) // 将字符串转为数组，以空格分隔。
	fmt.Printf("%s\t%s\n", reflect.TypeOf(strSli), strSli)
	for _, val := range strSli {
		fmt.Printf("%s ", val)
	}

	fmt.Println()
	str2 := strings.Join(strSli, ";")
	fmt.Printf("%s\n", str2)
}
