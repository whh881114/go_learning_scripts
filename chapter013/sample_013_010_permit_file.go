package main

import (
	"log"
	"os"
)

var (
	file     *os.File
	err      error
	dirpath  = "/tmp/test/test2/"
	fileName = "demo"
	filePath = dirpath + fileName
)

func main() {
	// 测试写权限
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Fatal("错误：没有写入权限。")
	} else if os.IsNotExist(err) {
		log.Fatal("错误：文件不存在。")
	} else {
		log.Printf("正确：文件%s有写入权限。", file.Name())
	}
	file.Close()

	// 测试读权限
	file, err = os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil && os.IsPermission(err) {
		log.Fatal("错误：没有读取权限。")
	} else if os.IsNotExist(err) {
		log.Fatal("错误：文件不存在。")
	} else {
		log.Printf("正确：文件%s有读取权限。", file.Name())
	}
	file.Close()
}
