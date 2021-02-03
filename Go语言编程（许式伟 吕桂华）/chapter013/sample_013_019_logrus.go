package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	// 只记录严重或以上的警告
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"tool":  "pen",
		"price": 10,
	}).Info("The pen price is 10 dlooars.")

	// }).Warn("The pen price is 10 dlooars.")
	// }).Fatal("The pen price is 10 dlooars.")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "这是一个字段",
		"other":  "其他你想记录的东西",
	})
	contextLogger.Info("此处会记录common和other字段。")
}
