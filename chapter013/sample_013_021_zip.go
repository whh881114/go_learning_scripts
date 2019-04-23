package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	// 创建一个打包文件
	outFile, err := os.Create("D:\\GoProjects\\src\\goTraining\\chapter013\\test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// 使用zip包的创建函数zipWriter用于写入文件
	zipWriter := zip.NewWriter(outFile)

	// 往打包文件中写文件
	// 这里使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入这个打包文件中。
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String content of file."},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	// 下面将要打包的内容写入打包文件中，依次写入
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = fileWriter.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 清理
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}
