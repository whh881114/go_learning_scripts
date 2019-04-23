package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	// 打开一个gzip文件
	// 文件是一个reader，但是可以使用各种数据源，比如Web服务器返回的gzipped内容，
	// 它的内容不是一个文件，而是一个内存流。
	gzipFile, err := os.Open("D:\\GoProjects\\src\\goTraining\\chapter013\\test3.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	defer gzipFile.Close()

	// 解压缩到一个writer，它是一个file writer
	outFileWriter, err := os.Create("D:\\GoProjects\\src\\goTraining\\chapter013\\test3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outFileWriter.Close()

	// 复制内容
	_, err = io.Copy(outFileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
