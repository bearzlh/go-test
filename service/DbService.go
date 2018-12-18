package service

import (
	"github.com/gohouse/gorose"
	"strconv"
)

type DbService struct {

}

var DB *gorose.Session

var Clusters = map[string]gorose.DbConfigCluster{}
var Sessions = map[string]*gorose.Session{}


//初始化用户连接
func GetUserDb(index int) *gorose.Session {
	name := "user"
	clusterIndex := name + strconv.Itoa(index % Cf.DB.Mysql.User.Total)

	if exists, ok := Sessions[clusterIndex]; ok {
		return exists
	}

	connection, err := gorose.Open(GetSpreadCluster("user", index, Cf.DB.Mysql.User))
	L.FailOnError(err, "连接失败")
	Sessions[name] = connection.NewSession()

	return Sessions[name]
}

//获取用户链接
func GetSpreadCluster(name string, userId int, mysql MysqlSpread) *gorose.DbConfigCluster {
	defaultParams := Cf.DB.Mysql.Default
	index := userId % mysql.Total
	clusterIndex := name + strconv.Itoa(int(index))
	if exists, ok := Clusters[clusterIndex]; ok {
		return &exists
	}

	dbName := mysql.Prefix + strconv.Itoa(int(index))
	cluster := gorose.DbConfigCluster{}
	for _, connection := range mysql.Connection {
		if index > connection.From && index < connection.To {
			for i := 0; i < len(connection.Read); i++ {
				hosts := connection.Read[i]
				slave := gorose.DbConfigSingle{}
				slave.Driver = Cf.DB.Driver
				slave.EnableQueryLog = true
				slave.Prefix = Cf.DB.Prefix
				slave.SetMaxIdleConns = 0
				slave.SetMaxOpenConns = 1
				slave.Dsn = defaultParams.Username + ":" + defaultParams.Password + "@tcp(" + hosts.Host + ":" + hosts.Port + ")/" + dbName
				cluster.Slave = append(cluster.Slave, &slave)
			}
			master := gorose.DbConfigSingle{}
			master.Driver = Cf.DB.Driver
			master.EnableQueryLog = true
			master.Prefix = Cf.DB.Prefix
			master.SetMaxIdleConns = 0
			master.SetMaxOpenConns = 1
			master.Dsn = defaultParams.Username + ":" + defaultParams.Password + "@tcp(" + connection.Write.Host + ":" + connection.Write.Port + ")/" + dbName
			cluster.Master = &master
			Clusters[clusterIndex] = cluster
		}
	}

	return &cluster
}

//初始化
func GetDefaultDb() *gorose.Session {
	name := "default"
	if exists, ok := Sessions[name]; ok {
		return exists
	}

	connection, err := gorose.Open(GetDefaultCluster(name, Cf.DB.Mysql.Default))
	L.FailOnError(err, "连接失败")
	Sessions[name] = connection.NewSession()

	return Sessions[name]
}

func GetDefaultCluster(name string,mysql Mysql) *gorose.DbConfigCluster {
	if exists, ok := Clusters[name]; ok {
		return &exists
	}
	cluster := gorose.DbConfigCluster{}
	for i := 0; i < len(mysql.Read); i++ {
		index := i
		slave := gorose.DbConfigSingle{}
		slave.Driver = Cf.DB.Driver
		slave.EnableQueryLog = true
		slave.Prefix = Cf.DB.Prefix
		slave.SetMaxIdleConns = 0
		slave.SetMaxOpenConns = 1
		slave.Dsn = mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Read[index].Host + ":" + mysql.Read[index].Port + ")/" + mysql.Database
		cluster.Slave = append(cluster.Slave, &slave)
	}
	master := gorose.DbConfigSingle{}
	master.Driver = Cf.DB.Driver
	master.EnableQueryLog = true
	master.Prefix = Cf.DB.Prefix
	master.SetMaxIdleConns = 0
	master.SetMaxOpenConns = 1
	master.Dsn = mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Write.Host + ":" + mysql.Write.Port + ")/" + mysql.Database
	cluster.Master = &master
	Clusters[mysql.Name] = cluster

	return &cluster
}