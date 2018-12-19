package test

import (
	"mq/service"
	"testing"
)

const envFile = "../.env"

func TestMain(m *testing.M) {
	service.GetConfig(envFile)
	service.GetLog(service.Cf.LogDir)
	m.Run()
}