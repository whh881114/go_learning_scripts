package mylog

// 自定义一个日志库，实现日志记录的功能

// 日志级别

// 定义日志级别
const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	CIRTAL
	FATAL
)

func getLevelStr(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "CIRTAL"
	case 6:
		return "FATAL"
	default:
		return "DEBUG"
	}
}
