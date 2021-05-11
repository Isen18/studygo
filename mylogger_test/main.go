package main

import (
	"time"

	"github.com/Isen18/studygo/mylogger"
)

var log mylogger.Logger

func main() {
	// log = mylogger.NewConsoleLogger("error")
	log = mylogger.NewFileLogger("Debug", "./", "isen.log", 10*1024)
	for i := 0; i < 1000; i++ {
		// for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warnning("这是一条warnning日志")
		log.Error("这是一条error日志, err=%v", "未知错误")
		log.Fatal("这是一条fatal日志")
	}

	time.Sleep(time.Second * 10)
}
