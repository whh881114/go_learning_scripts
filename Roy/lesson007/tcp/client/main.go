package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		fmt.Println("连接服务端失败，err：", err)
		return
	}

	defer conn.Close()

	// fmt.Fprintln(conn, "你好！")
	// _, err = conn.Write([]byte("你好！"))

	var input string
	fmt.Scan(&input)
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("发送消息失败，err：", err)
		return
	}
}
