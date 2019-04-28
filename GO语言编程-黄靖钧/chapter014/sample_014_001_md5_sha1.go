package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi, Golang!"
	Md5Inst := md5.New()              // 初始化md5实例
	Md5Inst.Write([]byte(TestString)) // 传入要计算的字符串
	Result := Md5Inst.Sum([]byte("")) // 开始计算md5值
	fmt.Printf("%x\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x", Result)
}
