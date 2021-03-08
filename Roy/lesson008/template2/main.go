package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 模板嵌套
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html", "ul.html")
	if err != nil {
		fmt.Printf("Parse failed: %s\n", err)
		return
	}

	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
