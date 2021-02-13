package main

import (
	"fmt"
	"net"
)

// 1. 临听端口
// 2. 接收客户端请求建立连接
// 3. 创建goroutine处理连接

func process(conn net.Conn) {
	// 单独的goroutine结束之后关闭连接
	defer conn.Close()

	// 从连接中接收数据
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("接收客户端发来的消息失败了，err：", err)
		return
	}
	fmt.Println("接收客户端发来的消息：", string(buf[:n]))
}

func main() {
	listner, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(conn)
	}
}
