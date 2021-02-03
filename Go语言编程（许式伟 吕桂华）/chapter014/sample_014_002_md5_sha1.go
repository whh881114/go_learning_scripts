package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "/Users/Kanon/go/src/goTraining/chapter014/sample_014_001_md5_sha1.go"
	infile, inerr := os.Open(fileName)
	if inerr != nil {
		fmt.Println(inerr)
		os.Exit(1)
	} else {
		md5h := md5.New()
		io.Copy(md5h, infile) // 计算文件摘要时，需要使用到io.Copy()函数。
		fmt.Printf("文件\"%s\"的md5值为：%x。\n", fileName, md5.Sum([]byte("")))

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("文件\"%s\"的sha1值为：%x。\n", fileName, sha1.Sum([]byte("")))
	}
}
