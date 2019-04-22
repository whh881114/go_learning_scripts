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
	// 可写方式打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将字节写入文件中
	file.Write([]byte("写入字节。\r\n"))
	// 将字符串写入文件中
	file.WriteString("写入字符串。\r\n")

	// 打印文件内容
	file, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	file.Close()
}
