package test

import (
	"fmt"
	"mq/service"
	"testing"
)

func TestGetSet(t *testing.T) {
	conn := service.GetRedis("default")
	getResult, err := conn.Do("SET", "A", "B")
	if err == nil {
		t.Log(getResult)
	}
	replyResult, err := conn.Do("GET", "C")
	if err != nil {
		t.Log(fmt.Sprintf("%s", replyResult))
	} else {
		t.Log("GET C error")
	}
}
