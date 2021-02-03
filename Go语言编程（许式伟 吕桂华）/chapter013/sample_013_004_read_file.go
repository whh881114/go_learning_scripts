package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile  *os.File
	fileInfo os.FileInfo
	err      error
	dirPath  = "/tmp/test/test2/"
	fileName = "demo"
	filePath = dirPath + fileName
)

func main() {
	// 简单地以只读的方式打开（读写的例子稍后再说）。
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 打印文件内容
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()

	fmt.Println("----------------------------------------------------------------------------------------")

	// OpenFile提供更多的选项，第一个参数是文件路径，第二个参数是打开时的属性，第三个参数是打开时的文件权限模式。
	// 第二个参数的属性可以单独使用，也可以组合使用，组合使用时可以使用OR操作符，例如：
	// os.O_CREATE|os.O_APPEND
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
}
