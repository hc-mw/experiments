package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rabbitmq/amqp091-go"
)

type Credentials struct {
	User string
	Pass string
	Host string
	Port string
}

type RabbitMQConn struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   amqp091.Queue
}

func NewRMQConn(c Credentials) RabbitMQConn {
	connObj := RabbitMQConn{}
	connObj.ConnectRMQ(c.User, c.Pass, c.Host, c.Port)
	return connObj
}

func (rm *RabbitMQConn) ConnectRMQ(user, pass, host, port string) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	conn, err := amqp091.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	rm.conn = conn

	ch, err := conn.Channel()
	FailOnError(err, "failed to open channel")
	rm.channel = ch
}

func (rm *RabbitMQConn) CloseConnection() {
	if rm.conn == nil {
		log.Println("cannot close nil connection")
		return
	}
	err := rm.channel.Close()
	FailOnError(err, "failed to close channel")

	err = rm.conn.Close()
	FailOnError(err, "failed to close channel connection")
}

func (rm *RabbitMQConn) GetConn() *amqp091.Connection {
	return rm.conn
}

func (rm *RabbitMQConn) GetChannel() *amqp091.Channel {
	return rm.channel
}

func (rm *RabbitMQConn) DeclareQueue(name string) amqp091.Queue {
	q, err := rm.channel.QueueDeclare("hello", false, false, false, false, nil)
	FailOnError(err, "failed to declare a queue")

	rm.queue = q
	return rm.queue
}

func (rm *RabbitMQConn) GetQueue() amqp091.Queue {
	return rm.queue
}

func (rm *RabbitMQConn) PublishOnQueue(ctx context.Context, queueName, message string) {
	err := rm.channel.PublishWithContext(ctx, "", queueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	FailOnError(err, "failed to publish the message")
	log.Printf(" [x] Sent %q\n", message)
}

func (rm *RabbitMQConn) ConsumeFromQueue(queueName string) <-chan amqp091.Delivery {
	msgs, err := rm.channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		FailOnError(err, "failed to consume the message")
	}
	return msgs
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
