package mylog

import (
	"fmt"
	"os"
	"time"
)

// FileLogger 往文件中记录日志的结构体
type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File // os包中File类型的指针
}

// NewFileLogger 是一个生成日志结构体实例
func NewFileLogger(level int, logFilePath string, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()
	return flObj
}

// 初始化日志文件
func (f *FileLogger) initFileLogger() {
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Open file: %s failed.", filepath))
	}
	f.logFile = file
}

// Debug 记录日志
// func (f *FileLogger) Debug(msg string) {
// 	f.logFile.WriteString(msg)
// }

// Debug 记录日志，第二版
func (f *FileLogger) Debug(format string, args ...interface{}) {
	// fmt.Fprintf(f.logFile, format, args...)
	// fmt.Fprintln(f.logFile)

	if f.level > DEBUG {
		return
	}

	fileName, funcName, line := getCallerInfo()

	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s\n", nowStr, getLevelStr(f.level), fileName, funcName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
}
