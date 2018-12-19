package service

import "github.com/gomodule/redigo/redis"

type RedisService struct {
}

var RedisClusters = map[string]redis.Conn{}

func GetRedis(name string) redis.Conn {
	if exists, ok := RedisClusters[name]; ok {
		return exists
	}

	var Config RedisConnection
	for _, config := range Cf.Cache.Connection {
		if config.Name == name {
			Config = config
			break
		}
	}

	option := redis.DialPassword(Config.Pass)
	conn, err := redis.Dial("tcp", Config.Host+":"+Config.Port, option)
	if err != nil {
		L.Debug("redis连接失败", "error")
	}

	RedisClusters[name] = conn
	return conn
}
