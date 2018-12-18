package test

import (
	"mq/service"
	"testing"
)

func TestLog(t *testing.T) {
	service.L.Debug("aadd", "error")
}