package service

import "github.com/gohouse/gorose"

type DbService struct {

}

var DB *gorose.Session

//保证单例模式
var DbBool = make(chan bool, 1)

//初始化
func GetDb() *gorose.Session {
	DbBool<-true
	if DB == nil {
		var DbConfig = Cf.DB
		connection, err := gorose.Open(&DbConfig)
		L.FailOnError(err, "连接失败")
		DB = connection.NewSession()
	}
	<-DbBool
	return DB
}
