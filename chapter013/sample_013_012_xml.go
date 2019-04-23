package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"` // 后面这个称为struct tag，用来辅助反射的。
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("D:\\GoProjects\\src\\goTraining\\chapter013\\sample_013_012_xml_servers.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(v)
}
