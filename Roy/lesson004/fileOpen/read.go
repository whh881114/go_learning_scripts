package main

import (
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

	var tmp [128]byte // 一次性读128个字节
	for {
		_, err := f.Read(tmp[:]) // 将读取的内容存入到tmp数组中。放弃此方法。
		if err == io.EOF {
			fmt.Println("文件读完了。")
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(tmp[:]))
	}
}
