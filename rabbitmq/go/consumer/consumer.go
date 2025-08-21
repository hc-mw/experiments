package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp091.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "failed to conenct to rabbitmq")

	ch, err := conn.Channel()
	failOnError(err, "failed to get a channel")

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "failed to declare queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("recieved a message: %q\n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
