package main

import "github.com/whh881114/go_learning_scripts/Roy/lesson004/log_demo/mylog"

// 写了一个项目，想要在代码中记录日志，需要使用mylog这个包。
func main() {
	f1 := mylog.NewFileLogger(mylog.DEBUG, "/tmp", "mylog_demo.log")
	f1.Debug("这是一条测试的日志。")

	userID := 10
	f1.Debug("id是%d的用户一直在尝试登陆。", userID)
}
