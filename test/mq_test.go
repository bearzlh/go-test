package test

import (
	"encoding/json"
	"fmt"
	"mq/helper"
	"mq/model"
	"mq/service"
	"os"
	"strconv"
	"testing"
	"time"
)

var L = service.LogService{}

const queueUserSubscribe = "collect.user.subscribe"
const queueReferralSubscribe = "collect.referral.subscribe"
const queueReferralPv = "collect.referral.pv"

type dotObject struct {
	UserId int64 `json:"user_id"`
	ChannelId int64 `json:"admin_id"`
	EventTime int64 `json:"event_time"`
	IsFirst uint8 `json:"is_first"`
	Sex string `json:"sex"`
	ReferralId int64 `json:"referral_id"`
	Type int `json:"type"`
}

func TestConsume(t *testing.T) {
	//初始化mq
	mq := service.GetMq()
	mq.GetQueue(mq.Channel, queueReferralSubscribe)
	start := time.Now().UnixNano()
	mq.Consume(queueReferralPv)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

/**
使用队列，初始化N个队列，2N个协程，N个协程向队列中添加数据，N个协程从队列中消费数据
10s     10000
    1   38401
    2   43197
    3   37485
    4   29187
    5   23074
 */
//使用队列处理数据
func TestUVWithChannel(t *testing.T) {
	//队列数量
	channelCount := 5
	//队列长度
	channelLen := 1000
	//执行时间
	timeLen := 10
	//吞吐量
	var total = 0

	var ReferralPvChannel = make([]chan string, channelCount)
	for next := 0; next < channelCount; next++ {
		ReferralPvChannel[next] = make(chan string, channelLen)
	}

	//初始化mq
	mq := service.GetMq()

	//初始化队列
	mq.GetQueue(mq.Channel, queueReferralPv)

	referralIds := model.GetReferralIds()

	//退出
	isFull := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(timeLen))
		isFull<-true
	}()

	//生产者发布
	for j := 0; j < channelCount; j++ {
		next := j
		go func() {
			for {
				select {
				case dot := <-ReferralPvChannel[next]:
					total++
					mq.Produce(dot, queueReferralPv)
				}
			}
		}()
	}

	for i := 0; i < channelCount; i++ {
		next := i
		go func() {
			ch1 := dotObject{};
			ch1.UserId = int64(helper.HA.RandInt(10000, 20000))
			rid := referralIds[helper.HA.RandInt(0, len(referralIds))]
			if rid.Valid {
				ch1.ReferralId = rid.Int64
			} else {
				ch1.ReferralId = 0
			}
			ch1.EventTime = time.Now().Unix()
			data,_ := json.Marshal(ch1)
			for ; len(isFull) == 0; {
				ReferralPvChannel[next]<-string(data)
			}
		}()
	}

	select {
		case <-isFull:
			fmt.Println("执行次数：" + strconv.Itoa(total))
			os.Exit(0)
	}
}

//不使用队列处理数据
func TestUVNoChannel(t *testing.T) {
	channelCount := 5

	//初始化mq
	mq := service.GetMq()

	//初始化队列
	mq.GetQueue(mq.Channel, queueReferralPv)

	referralIds := model.GetReferralIds()

	//退出
	isFull := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 10)
		isFull<-true
	}()

	for i := 0; i < channelCount; i++ {
		go func() {
			ch1 := dotObject{};
			ch1.UserId = int64(helper.HA.RandInt(10000, 20000))
			rid := referralIds[helper.HA.RandInt(0, len(referralIds))]
			if rid.Valid {
				ch1.ReferralId = rid.Int64
			} else {
				ch1.ReferralId = 0
			}
			ch1.EventTime = time.Now().Unix()
			data,_ := json.Marshal(ch1)
			for ; len(isFull) == 0; {
				mq.Produce(string(data), queueReferralPv)
			}
		}()
	}

	select {
	case <-isFull:
		os.Exit(0)
	}
}

//数据发送整体测试
func TestSend(t *testing.T) {
	var ReferralSubscribeChannel = make(chan string, 10000)
	var UserSubscribe = make(chan string, 10000)
	var ReferralPv = make(chan string, 10000)
	//初始化mq
	mq := service.GetMq()

	//初始化队列
	mq.GetQueue(mq.Channel, queueReferralSubscribe)
	mq.GetQueue(mq.Channel, queueUserSubscribe)
	mq.GetQueue(mq.Channel, queueReferralPv)

	referralIds := model.GetReferralIds()

	//退出
	isFull := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 10)
		//isFull<-true
	}()

	//生产者发布
	go func() {
		for {
			select {
			case dot := <-ReferralSubscribeChannel:
				mq.Produce(dot, queueReferralSubscribe)
			case dot := <-UserSubscribe:
				mq.Produce(dot, queueUserSubscribe)
			case dot := <-ReferralPv:
				mq.Produce(dot, queueReferralPv)
			}
		}
	}()

	//模糊关注
	//for i := 0; i < 10; i++ {
	go func() {
		ch1 := dotObject{};
		rid := referralIds[helper.HA.RandInt(0, len(referralIds))]
		if rid.Valid {
			ch1.ReferralId = rid.Int64
		} else {
			ch1.ReferralId = 0
		}
		ch1.Type = 4
		ch1.EventTime = time.Now().Unix()
		data,_ := json.Marshal(ch1)
		for ; len(isFull) == 0; {
			ReferralSubscribeChannel<-string(data)
		}
	}()
	//}

	//用户扫码关注
	//for i := 0; i < 10; i++ {
	go func() {
		ch1 := dotObject{};
		ch1.ChannelId = 3
		ch1.IsFirst = 0
		ch1.Type = 1
		ch1.EventTime = time.Now().Unix()
		data,_ := json.Marshal(ch1)
		for ; len(isFull) == 0; {
			UserSubscribe<-string(data)
		}
	}()
	//}

	//uv
	//for i := 0;i < 10;i++  {
	go func() {
		ch1 := dotObject{};
		ch1.UserId = int64(helper.HA.RandInt(10000, 20000))
		rid := referralIds[helper.HA.RandInt(0, len(referralIds))]
		if rid.Valid {
			ch1.ReferralId = rid.Int64
		} else {
			ch1.ReferralId = 0
		}
		ch1.EventTime = time.Now().Unix()
		data,_ := json.Marshal(ch1)
		for ; len(isFull) == 0; {
			ReferralPv<-string(data)
		}
	}()
	//}

	select {
	case <-isFull:
		os.Exit(0)
	}
}