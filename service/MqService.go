package service

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MqService struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
}

var Mq *MqService
var L = LogService{}

//保证单例模式
var ChBool = make(chan bool, 1)

//初始化
func GetMq() *MqService {
	ChBool<-true
	if Mq == nil {
		mqConfig := Cf.Mq

		Mq = &MqService{}
		Mq.Connection = getConnection(mqConfig.User, mqConfig.Passwd, mqConfig.Host, mqConfig.Port, mqConfig.Vhost)

		Ch, _ := Mq.Connection.Channel()

		Mq.Channel = Ch

		Mq.setExchange(Ch, mqConfig.Exchange)
	}
	<-ChBool

	return Mq
}

//获取连接
func getConnection(user string, passwd string, host string, port string, vhost string) *amqp.Connection {
	hostPath := "amqp://"+user+":"+passwd+"@"+host+":"+port+"/"+vhost
	L.Debug("connect==>" + hostPath, LEVEL_DEBUG)
	conn, err := amqp.Dial(hostPath)
	L.FailOnError(err, "Failed to connect to RabbitMQ")
	L.Debug("connected!", LEVEL_DEBUG)
	return conn
}

//声明队列
func (M *MqService)GetQueue(channel *amqp.Channel, channelName string) amqp.Queue {
	Q, err := channel.QueueDeclare(
		channelName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	L.FailOnError(err, "Failed to declare a queue")
	return Q
}

//声明交换机
func (M *MqService)setExchange(channel *amqp.Channel, exchangeName string)  {
	err := channel.ExchangeDeclare(
		exchangeName,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	L.FailOnError(err, "exchange get error")
}

//关闭连接
func (M *MqService)CloseMq() error {
	err := Mq.Connection.Close()
	if err != nil{
		return err
	}

	err1 := Mq.Channel.Close()
	if err != nil {
		return err1
	}

	return nil
}

//生产
func (M *MqService) Produce(body string, queueName string) {
	err := M.Channel.Publish(
		"",                    // exchange
		queueName, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	L.Debug(fmt.Sprintf("==>Sent %s:", body), LEVEL_DEBUG)
	if err != nil {
		fmt.Println(err)
	}
}

//消费
func (M *MqService) Consume(queueName string) {
	messages, _ := M.Channel.Consume(
		queueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range messages {
		L.Debug(fmt.Sprintf("<==Received: %s", d.Body), LEVEL_DEBUG)
	}
}