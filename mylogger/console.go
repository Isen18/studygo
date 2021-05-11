package mylogger

import (
	"fmt"
)

// ConsoleLogger 结构体
type ConsoleLogger struct {
	//日志级别
	level LogLevel
}

// ConsoleLogger 构造行数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return ConsoleLogger{
		level: level,
	}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.level
}

func (c ConsoleLogger) log(level LogLevel, msg string) {
	if c.enable(level) {
		fileName, funcName, lineNo := getCallerInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", getTimeStr(), getLogLevelStr(level), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, fmt.Sprintf(format, a...))
}

func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, fmt.Sprintf(format, a...))
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, fmt.Sprintf(format, a...))
}

func (c ConsoleLogger) Warnning(format string, a ...interface{}) {
	c.log(WARNING, fmt.Sprintf(format, a...))
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, fmt.Sprintf(format, a...))
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, fmt.Sprintf(format, a...))
}
