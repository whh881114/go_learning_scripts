package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		return
	}

	daytime := time.Now().String()
	// conn.WriteToUDP([]byte(daytime), addr)

	// 自已添加的提示消息。
	pmt := fmt.Sprintf(" FROM: %s", addr)
	msg := strings.Join([]string{daytime, pmt}, "")
	conn.WriteToUDP([]byte(msg), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：", err.Error())
		os.Exit(1)
	}
}
