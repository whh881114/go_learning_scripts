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
	path     = "/tmp/test/test2/"
	fileName = "demo"
	filePath = path + fileName
)

func main() {
	file, _ := os.Open(filePath)
	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5

	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	whence := 0

	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("移动到位置5：", newPosition)

	// 从当前位置回退两字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("从肖前位置退回两字节：", newPosition)

	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("转到文件开始位置(0,0)：", newPosition)
}
