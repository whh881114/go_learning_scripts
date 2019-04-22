package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newfile  *os.File
	fileInfo os.FileInfo
	err      error
	path     = "/tmp/test/test2/"
	fileName = "demo"
	filePath = path + fileName
)

func main() {
	// 创建一个硬链接
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个
	// 删除和重命名不会影响一个
	hardLink := filePath + "hl"
	err := os.Link(filePath, hardLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建硬链接")

	// 创建一个软链接
	softLink := filePath + "sl"
	err = os.Symlink(fileName, softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建软链接")

	// lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	// Symlink在Windows中不工作
	fileInfo, err := os.Lstat(softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("链接信息：%+v", fileInfo)

	// 改变软链接的拥有者不会影响原始文件
	err = os.Lchown(softLink, os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
