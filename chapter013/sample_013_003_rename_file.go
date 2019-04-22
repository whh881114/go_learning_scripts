package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	originalPath := "/tmp/test.txt"
	newPath := "/tmp/test2.txt"

	newFile, err := os.Create(originalPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("测试文件:%s创建成功。\n", newFile.Name())
	}

	err = os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("文件重命名成功，原文件\"%s\"，新文件\"%s\"。", originalPath, newPath)
	}
}
