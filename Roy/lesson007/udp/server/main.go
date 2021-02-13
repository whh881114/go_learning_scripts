package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		// IP: net.IPv4(0,0,0,0),
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("启动server失败，err：", err)
		return
	}

	defer listener.Close()

	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败，err：", err)
			return
		}
		fmt.Printf("接收到来自客户端%v的消息：%v\n", addr, string(buf[:n]))

		n, err = listener.WriteToUDP([]byte("来自服务器端：滚，你丫的！"), addr)
		if err != nil {
			fmt.Println("回复消息失败，err：", err)
			return
		}
	}
}
