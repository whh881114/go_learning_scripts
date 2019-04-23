package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	outputFile, err := os.Create("D:\\GoProjects\\src\\goTraining\\chapter013\\test3.txt.gz")
	if err != nil {
		log.Fatal(err)
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// 当写入gzip writer数据时，分它依次压缩数据并写入底层的文件中
	// 不必关心它是如何压缩的，还是像普通的writer一样操作即可。
	_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("已经压缩数据并写入文件。")
}
