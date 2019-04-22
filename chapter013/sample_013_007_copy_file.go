package main

import (
	"io"
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
	// 打开原始文件
	originalFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// 创建新的文件作为目标文件
	newFile, err := os.Create(filePath + "_copy")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// 从源文件中复制字切到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("文件已复制，大小 %d bytes。", bytesWritten)

	// 将文件内容flush到硬盘中。
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

// 复制操作有两点需要注意的地方：
// 第一点是Create()执行之后需要Close()关闭回收资源。
// 第二点是调用io包中的复制函数之后文件内容并没有真正保存在文件中，使用Sync()函数同步之后才能真正保存了在硬盘中。
