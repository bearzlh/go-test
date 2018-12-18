package test

import (
	"mq/model"
	"mq/service"
	"testing"
)

func TestJson(t *testing.T) {
	t.Log(service.Cf.DB)
}

func TestGetDsn(t *testing.T) {
	user := model.User{}

	session := service.GetUserDb(1)
	session.Table(&user).First()
	t.Log(user.OpenId)
}