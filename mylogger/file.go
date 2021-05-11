package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"
)

var (
	logMsgQueueLimit = 50000
)

// FileLogger 结构体
type FileLogger struct {
	level       LogLevel //日志级别
	filePath    string   //日志文件路径
	fileName    string   //日志文件名称
	maxFileSize int64    //日志文件最大大小
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *LogMsg
}

// LogMsg 结构体
type LogMsg struct {
	level     LogLevel
	msg       string
	fileName  string
	lineNo    int
	funcName  string
	timestamp string
}

// FileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	log := &FileLogger{
		level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxFileSize,
		logChan:     make(chan *LogMsg, logMsgQueueLimit),
	}

	err = log.initFile()
	if err != nil {
		panic(err)
	}
	return log
}

//打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err=%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err=%v\n", err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errFileObj

	go f.writeLogBackground()
	return nil
}

// 关闭日志文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 是否需要写日志
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.level
}

// 是否需要切割日志文件
func (f *FileLogger) checkSplit(file *os.File) (bool, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat() failed, err=%v\n", err)
		return false, errors.New("获取文件信息出错")
	}

	return fileInfo.Size() > f.maxFileSize, nil
}

// 如果有必要，则进行日志切割
func (f *FileLogger) splitFileIfNeed(file *os.File) (*os.File, error) {
	ok, err := f.checkSplit(file)
	if err != nil {
		return file, err
	}

	if !ok {
		//无需切割日志文件
		return file, nil
	}

	//切割文件
	//1. 关闭文件
	file.Close()

	//2. 备份文件
	nowStr := time.Now().Format("20160102150405000")
	logName := path.Join(f.filePath, file.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(logName, newLogName)

	//3. 打开新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return fileObj, err
}

// 后台写日志到文件中
func (f *FileLogger) writeLogBackground() {
	for {
		fileObj, err := f.splitFileIfNeed(f.fileObj)
		if err != nil {
			fmt.Printf("f.splitFileIfNeed(f.fileObj) failed, err=%v\n", err)
			return
		}
		f.fileObj = fileObj

		// select {
		// case logMsg := <-f.logChan: //这个阻塞会消费cpu?
		// 	log := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", getTimeStr(), getLogLevelStr(logMsg.level), logMsg.fileName, logMsg.funcName, logMsg.lineNo, logMsg.msg)

		// 	fmt.Fprint(fileObj, log)

		// 	if logMsg.level >= ERROR {
		// 		errFileObj, err := f.splitFileIfNeed(f.errFileObj)
		// 		if err != nil {
		// 			fmt.Printf("f.splitFileIfNeed(f.errFileObj) failed, err=%v\n", err)
		// 			return
		// 		}

		// 		f.errFileObj = errFileObj
		// 		fmt.Fprint(errFileObj, log)
		// 	}
		// default:
		// 	//没有日志休息100ms
		// 	fmt.Println("没有日志休息100ms")
		// 	time.Sleep(time.Microsecond * 100)
		// }

		logMsg := <-f.logChan
		log := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", getTimeStr(), getLogLevelStr(logMsg.level), logMsg.fileName, logMsg.funcName, logMsg.lineNo, logMsg.msg)

		fmt.Fprint(fileObj, log)

		if logMsg.level >= ERROR {
			errFileObj, err := f.splitFileIfNeed(f.errFileObj)
			if err != nil {
				fmt.Printf("f.splitFileIfNeed(f.errFileObj) failed, err=%v\n", err)
				return
			}

			f.errFileObj = errFileObj
			fmt.Fprint(errFileObj, log)
		}
	}
}

func (f *FileLogger) log(level LogLevel, msg string) {
	if f.enable(level) {
		fileName, funcName, lineNo := getCallerInfo(3)

		logMsg := &LogMsg{
			level:     level,
			msg:       msg,
			fileName:  fileName,
			lineNo:    lineNo,
			funcName:  funcName,
			timestamp: getTimeStr(),
		}

		select {
		case f.logChan <- logMsg:
		default: // 通道满了，则丢弃日志
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, fmt.Sprintf(format, a...))
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, fmt.Sprintf(format, a...))
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, fmt.Sprintf(format, a...))
}

func (f *FileLogger) Warnning(format string, a ...interface{}) {
	f.log(WARNING, fmt.Sprintf(format, a...))
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, fmt.Sprintf(format, a...))
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, fmt.Sprintf(format, a...))
}
