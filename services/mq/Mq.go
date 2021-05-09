package mq

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://admin:admin@10.211.55.6:5672/")
	return conn, err
}

// Publish 发送端函数
func Publish(exchange, queueName, body string) error {
	//建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	//创建通道channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	//创建队列
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//发送消息
	channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return err
}

// Consumer 接收者方法
func Consumer(exchange, queueName string, callback Callback) {
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建通道
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建queue
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Println("waiting for message")
	<-forever

}
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func PublishEx(exchange, types, routingKey, body string) error {
	//建立连接
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		return err
	}
	//创建通道
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		return err
	}
	//创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return err
}

func ConsumerEx(exchange, types, routingKey string, callback Callback) {
	//建立连接
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建通道
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建交换机
	err = channel.ExchangeDeclare(
		exchange,
		types,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	q, err := channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	//绑定
	err = channel.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := channel.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Printf("waiting for message\n")
	<-forever
}

// ConsumerDlx 死信队列消费端
func ConsumerDlx(exchangeA, queueAName, exchangeB, queueBName string, ttl int, callback Callback) {
	//建立连接
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建channel
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建A交换机
	//创建A队列
	//A交换机和A队列绑定
	err = channel.ExchangeDeclare(
		exchangeA,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	queueA, err := channel.QueueDeclare(
		queueAName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-message-ttl":          ttl,
			"x-dead-letter-exchange": exchangeB,
			//"x-dead-letter-queue":"",
			//"x-dead-letter-routing-key":"",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = channel.QueueBind(
		queueA.Name,
		"",
		exchangeA,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建B交换机
	//创建B队列
	//B交换机和B队列绑定
	err = channel.ExchangeDeclare(
		exchangeB,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	queueB, err := channel.QueueDeclare(
		queueBName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = channel.QueueBind(
		queueB.Name,
		"",
		exchangeB,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgs, err := channel.Consume(queueB.Name, "", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			callback(*s)
			d.Ack(false)
		}
	}()
	fmt.Println("Waiting for messages")
	<-forever
}

// PublishDlx 死信队列生产端
func PublishDlx(exchangeA, body string) error {
	//建立连接
	conn, err := Connect()
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//创建channel
	channel, err := conn.Channel()
	defer channel.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	return err
}
