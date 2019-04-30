package main

import "fmt"

func main() {
	var a *int // 定义一个a指针，指向nil
	aP := &a   // 空指针a的指针

	fmt.Printf("a--->nil: %x\n", a)
	fmt.Printf("aP--->a: %x\n", aP)
	fmt.Printf("aP--->a--->nil（指针aP指向的指针a的内存地址）：%x\n", *aP)
	fmt.Printf("&aP--->aP（表示aP在内存中的地址）：%x\n", &aP)

	// 定义一个指针的指针，可以这样定义：var ptr **int。
	fmt.Println("*****************************************")
	b := 10
	bP := &b
	bPP := &bP
	fmt.Printf("b: %d\n", b)
	fmt.Printf("bP: %d\n", bP)
	fmt.Printf("*bP: %d\n", *bP)
	fmt.Printf("bPP: %d\n", bPP)
	fmt.Printf("*bPP: %d\n", *bPP)
	fmt.Printf("**bPP: %d\n", **bPP)
}
