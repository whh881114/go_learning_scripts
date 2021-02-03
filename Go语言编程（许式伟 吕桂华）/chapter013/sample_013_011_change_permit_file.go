package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	file     *os.File
	err      error
	dirpath  = "/tmp/test/test2/"
	fileName = "demo"
	filePath = dirpath + fileName
)

func main() {
	// 使用Linux风格改变文件权限
	err := os.Chmod(filePath, 0777)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("修改文件权限成功。")
	}

	// 改变文件所有者
	err = os.Chown(filePath, os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("修改文件所有者成功。")
	}

	// 查看文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("文件不存在。")
		}
		log.Fatal(err)
	}
	fmt.Println("最白后修改时间：", fileInfo.ModTime())

	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes(filePath, lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
