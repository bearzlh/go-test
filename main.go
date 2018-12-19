package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"mq/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const queueNme = "test"
const envFile = "./.env"

type dotObject struct {
	EventTime string `json:"event_time"`
	ReferralId string `json:"referral_id"`
	DotType string `json:"type"`
}

var L = service.LogService{}

func main() {
	//初始化配置
	service.GetConfig(envFile)

	ReferralChannel := make(chan string, 1000)

	L := service.GetLog(service.Cf.LogDir)
	//初始化mq
	service.GetMq()

	mq := service.Mq
	mq.GetQueue(mq.Channel, queueNme)

	//收到终止信号的处理
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		L.Debug("Closing Mq...", service.LEVEL_DEBUG)
		if err := mq.CloseMq();err != nil {
			L.Debug("close mq error", service.LEVEL_ERROR)
		}
		os.Exit(0)
	}()

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	//生产者发布
	go func() {
		for {
			select {
			case dot := <-ReferralChannel:
				mq.Produce(dot, queueNme)
			}
		}
	}()

	isFull := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second)
		isFull<-true
	}()

	//生产者队列
	go func() {
		ch1 := dotObject{};
		u := uuid.Must(uuid.NewV4())
		ch1.ReferralId = u.String()
		ch1.DotType = "4"
		ch1.EventTime = time.Now().Format("2006/01/02 15:04:05")
		data,_ := json.Marshal(ch1)
		for ; len(isFull) == 0; {
			ReferralChannel<-string(data)
		}
	}()

	//消费数据
	go func() {
		mq.Consume(queueNme)
	}()

	L.Debug(fmt.Sprint("...Waiting for messages. To exit press CTRL+C"), service.LEVEL_DEBUG)

	//阻塞主进程，避免退出
	select {}
}