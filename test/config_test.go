package test

import (
	"mq/service"
	"testing"
)

func TestConfig(t *testing.T) {
	service.GetConfig().LoadFile("/Users/Bear/gopath/src/mq/config/mq.json")
	t.Log(service.GetConfig().Mq.Exchange)
}
