package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
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

// 所有的key都是大写，这样在现实的中肯定不会顺序，这个时候需要使用struct tag来定义。
