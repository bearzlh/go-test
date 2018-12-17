package test

import (
	"mq/helpter"
	"mq/model"
	"mq/service"
	"testing"
)

var AH = helpter.ArrayHelper{}

func TestJson(t *testing.T) {
	service.GetConfig().LoadFile("/Users/Bear/gopath/src/mq/config/config.json")
	referralIds := model.GetReferralIds()
	t.Log(AH.RandInt(0, len(referralIds)))
}