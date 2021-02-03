package main

import "fmt"

type Camera struct{}

func (c *Camera) TakeAPicture() string { return "拍照" }

type Phone struct{}

func (p *Phone) Call() string { return "响铃" }

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("我们的新款拍照手机有多种功能：")
	fmt.Println("打开了相机：", cp.TakeAPicture())
	fmt.Println("电话来电：", cp.Call())
}
