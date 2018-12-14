package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"mq/mq"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type dotObject struct {
	EventTime string `json:"event_time"`
	ReferralId string `json:"referral_id"`
	DotType string `json:"type"`
}

func main() {
	ReferralChannel := make(chan dotObject, 10)

	var stopLock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		stopLock.Lock()
		stop = true
		stopLock.Unlock()
		log.Println("Close Mq...")
		err := mq.CloseMq()
		if err != nil {
			fmt.Println("close mq error" + err.Error())
		}
		stopChan <- struct{}{}
		os.Exit(0)
	}()

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case dot := <-ReferralChannel:
				str,_ := json.Marshal(dot)
				body := string(str)
				err := mq.GetMq().Channel.Publish(
					"",                    // exchange
					mq.GetMq().Queue.Name, // routing key
					false,                 // mandatory
					false,                 // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				log.Printf(" [x] Sent %s", body)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	go func() {
		ch1 := dotObject{};
		ch1.ReferralId = "706"
		ch1.DotType = "4"
		ch1.EventTime = "1544611770"
		for {
			ReferralChannel<-ch1
			time.Sleep(time.Second * 1)
		}
	}()

	mq.Consume()

	select {
	}
}