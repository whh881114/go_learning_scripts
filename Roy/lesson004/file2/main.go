package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("/Users/roy/bin/auto_scaling_groups.sh")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close() // 使用defer延迟头闭文件

	// 利用缓冲区从文件读取
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(str)
	}
}
