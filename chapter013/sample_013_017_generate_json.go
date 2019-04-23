package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"server_name"`
	ServerIP   string `json:"server_ip"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Local_Web", ServerIP: "172.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Local_DB", ServerIP: "172.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err: ", err)
	}
	fmt.Println(string(b))
}

// 使用struct tag来定义不同的key名字。
