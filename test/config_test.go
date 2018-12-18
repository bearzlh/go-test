package test

import (
	"mq/service"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log(service.Cf.DB)
}
