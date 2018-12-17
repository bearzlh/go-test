package test

import (
	"mq/service"
	"testing"
)

func TestConfig(t *testing.T) {
	service.GetConfig().LoadFile("/Users/Bear/gopath/src/mq/config/config.json")
	t.Log(service.GetConfig().DB.Driver)
}
