package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

const user string = "rabbitmquser"
const passwd string = "rabbitmqpass"
const host string = "localhost"
const vhost string = ""
const exchange string = "cps.ex.topic"
const channel string = "collect.referral.subscribe"

type MqService struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Queue amqp.Queue
}

type dotObject struct {
	EventTime string `json:"event_time"`
	ReferralId string `json:"referral_id"`
	DotType string `json:"type"`
}

var Mq *MqService

var lock *sync.Mutex = &sync.Mutex {}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetMq() *MqService {
	lock.Lock()
	defer lock.Unlock()
	if Mq == nil {
		Mq = &MqService{}
		conn, err := amqp.Dial("amqp://"+user+":"+passwd+"@"+host+":5672/"+vhost)
		failOnError(err, "Failed to connect to RabbitMQ")

		Mq.Connection = conn

		Ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")

		Mq.Channel = Ch

		Ch.ExchangeDeclare(
			exchange,
			"topic",
			false,
			false,
			false,
			false,
			nil,
		)

		Q, err := Ch.QueueDeclare(
			channel, // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")

		Mq.Queue = Q
	}

	return Mq
}

func Publish(body dotObject){
	str,_ := json.Marshal(body)
	content := string(str)
	e := GetMq().Channel.Publish(
		"",     // exchange
		GetMq().Queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(content),
		})
	log.Printf(" [x] Sent %s", content)
	failOnError(e, "Failed to publish a message")
}

func Consume() {
	msgs, err := GetMq().Channel.Consume(
		GetMq().Queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func CloseMq() error {
	err := GetMq().Connection.Close()
	if err != nil{
		return err
	}

	err1 := GetMq().Channel.Close()
	if err != nil {
		return err1
	}

	return nil
}