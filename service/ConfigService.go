package service

import (
	"encoding/json"
	"fmt"
	"github.com/gohouse/gorose"
	"io/ioutil"
)

type ConfigService struct {
	Mq mq
	DB gorose.DbConfigSingle
	Debug bool
}

type mq struct {
	User string `json:"user"`
	Passwd string `json:"passwd"`
	Host string `json:"host"`
	Port string `json:"port"`
	Vhost string `json:"vhost"`
	Exchange string `json:"exchange"`
}

var Cf *ConfigService

//保证单例模式
var ConfigBool = make(chan bool, 1)

//初始化
func GetConfig() *ConfigService {
	ConfigBool<-true
	if Cf == nil {
		fmt.Println("----")
		Cf = &ConfigService{}
	}
	<-ConfigBool
	return Cf
}

//加载配置文件
func (C *ConfigService)LoadFile(file string) *ConfigService {
	L := LogService{}
	content, err :=ioutil.ReadFile(file)
	L.FailOnError(err, "文件内容读取失败")

	err = json.Unmarshal(content, &Cf)
	L.FailOnError(err, "内容解析错误")

	return C
}