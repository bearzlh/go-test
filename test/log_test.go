package test

import (
	"mq/service"
	"testing"
)

func TestLog(t *testing.T) {
	L := service.LogService{}
	for i := 0; i < 10; i++ {
		L.Debug("aadd", "error")
	}
}