package service

import (
	"fmt"
	"io"
	"os"
	"time"
)

const LEVEL_DEBUG = "debug";
const LEVEL_INFO = "info";
const LEVEL_NOTICE = "notice";
const LEVEL_ERROR = "error";
const LEVEL_ALERT = "alert";
const LEVEL_CRITICAL = "critical";

type LogService struct {
	LogDir string
}

var L *LogService

var chLog = make(chan bool, 1)

func GetLog(logDir string) *LogService {
	chLog <- true
	if L == nil {
		L = &LogService{}
		L.LogDir = logDir
	}
	<- chLog

	return L
}

//错误处理
func (Log *LogService)FailOnError(err error, msg string) {
	if err != nil {
		L.Debug(msg, LEVEL_ERROR)
		L.outPut(fmt.Sprintf("%s: %s\n", msg, err))
		os.Exit(1)
	}
}

func (Log *LogService)DebugOnError(err error) {
	if err != nil {
		L.Debug(err.Error(), LEVEL_DEBUG)
		L.outPut(fmt.Sprintf("%s\n", err))
	}
}

//打日志
func (Log *LogService) Debug(msg string, level string) {
	if level == "" {
		level = LEVEL_DEBUG
	}
	now := time.Unix(time.Now().Unix(), 0)
	dir := L.LogDir + "/"+fmt.Sprintf("%d%d",now.Year(), now.Month())

	//检查目录
	fileInfo, err := os.Stat(dir)
	if fileInfo == nil || !fileInfo.IsDir() {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			L.outPut(fmt.Sprintf("%s\n", err))
			return
		}
	}

	logFile := dir+"/"+fmt.Sprintf("%d_%d",now.Day(), now.Hour())+".log"
	file, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		L.outPut(fmt.Sprintf("%s\n", err))
		return
	}
	if Cf.Debug {
		_, err1:= io.WriteString(file, L.getMsg(msg, level));
		if err != nil {
			L.outPut(fmt.Sprintf("%s\n", err1))
		}
	}
}

//日志信息格式化
func (Log *LogService) getMsg(msg string, level string) string {
	Nano := time.Now().Nanosecond() / 1000000
	msg = fmt.Sprintf("level:[%s]\ttime:[%s.%d]\tpid:[%d]\tmsg:[%s]", level, time.Now().Format("2006-01-02 15:04:05"), Nano, os.Getpid(), msg)
	L.outPut(msg)
	return msg + "\n"
}

func (Log *LogService)outPut(msg string) {
	if Cf.Debug {
		fmt.Println(msg)
	}
}