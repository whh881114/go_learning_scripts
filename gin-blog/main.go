package main

import (
	"fmt"
	"net/http"
)
import "github.com/whh881114/go_learning_scripts/gin-blog/pkg/setting"
import "github.com/whh881114/go_learning_scripts/gin-blog/routers"

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
