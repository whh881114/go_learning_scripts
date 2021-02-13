package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("连接server失败，err：", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("滚！"))
	if err != nil {
		fmt.Println("发送消息失败，err：", err)
		return
	}
	var buf [1024]byte
	_, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("读取消息失败，err：", err)
		return
	}
	fmt.Println("收到回复：", string(buf[:]))
}
