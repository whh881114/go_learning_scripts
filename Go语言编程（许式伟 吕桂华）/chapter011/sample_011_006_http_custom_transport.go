package main

import (
	"crypto/tls"
	"net/http"
)

func main() {
	// 自定义http.Transport，就也就是重定Transport结构体里的参数而已。
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
		DisableCompression: true
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://example.com")
}
