package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const ConfigName = "default"
const Ext = ".json"
const ConfigDir = "config"
const Split = "/"

type ConfigService struct {
	Env        string
	ConfigPath string
	Debug      bool
	Mq         mq
	DB         db
	LogDir 	   string
}

//数据库全局配置
type db struct{
	Driver string `json:"driver"`
	Prefix string `json:"prefix"`
	Charset string `json:"charset"`
	Mysql struct{
		Default Mysql
		User MysqlSpread
	}
}

type MysqlSpread struct {
	Prefix string `json:"prefix"`
	Total int `json:"total"`
	Connection []struct{
		From int `json:"from"`
		To int `json:"to"`
		Read []hostPort
		Write hostPort
	}
}

//数据库基本配置
type Mysql struct {
	Name string `json:"name"`
	Charset string `json:"charset"`
	Collation string `json:"collation"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Prefix  string `json:"prefix"`
	Read []hostPort
	Write hostPort
}

//IP端口
type hostPort struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

//mq配置
type mq struct {
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Vhost    string `json:"vhost"`
	Exchange string `json:"exchange"`
}

var Cf *ConfigService

//保证单例模式
var ConfigBool = make(chan bool, 1)

//初始化配置，加载config配置文件
func GetConfig(env string) *ConfigService {
	ConfigBool <- true
	if Cf == nil {
		Cf = &ConfigService{}
		Cf.Env, Cf.ConfigPath = checkEnv(env)
		Cf.Debug = true
		Cf.loadFile()
	}
	<-ConfigBool
	return Cf
}

func checkEnv(envFile string) (string, string) {
	env := "default"

	file, _ := filepath.Abs(envFile)

	fileState, err := os.Stat(envFile)

	if err != nil {
		fmt.Println("配置文件错误")
		os.Exit(1)
	} else {
		if !fileState.IsDir() {
			content, _ := ioutil.ReadFile(envFile)
			env = string(content)
		}
	}

	dir := filepath.Dir(file)
	return env, dir
}

//加载配置文件
func (C *ConfigService) loadFile() *ConfigService {
	defaultFile := Cf.ConfigPath + Split + ConfigDir + Split + ConfigName + Ext
	envFile := Cf.ConfigPath + Split + ConfigDir + Split + Cf.Env + Ext
	L := LogService{}
	content, err := ioutil.ReadFile(defaultFile)
	L.FailOnError(err, "默认文件内容读取失败")

	err = json.Unmarshal(content, &Cf)
	L.FailOnError(err, "默认内容解析错误")

	envContent, err := ioutil.ReadFile(envFile)
	L.DebugOnError(err)

	err = json.Unmarshal(envContent, &Cf)
	L.DebugOnError(err)
	return C
}
