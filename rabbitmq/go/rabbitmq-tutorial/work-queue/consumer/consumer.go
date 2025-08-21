package main

import (
	"bytes"
	"log"
	"rabbitmq-tut/utils"
	"sync/atomic"
	"time"
)

var id atomic.Uint64

func init() {
	id.Store(1)
	log.Println("init method")
}

func main() {
	connObj := utils.NewRMQConn(utils.Credentials{
		User: "user",
		Pass: "password",
		Host: "localhost",
		Port: "5672",
	})
	defer connObj.CloseConnection()

	q := connObj.DeclareQueue("hello")

	msgs := connObj.ConsumeFromQueue(q.Name)

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			v := id.Load()
			log.Printf("recieved a message: %q, assigned id: %d\n", d.Body, v)

			dotCount := bytes.Count(d.Body, []byte("."))
			time.Sleep(time.Duration(dotCount) * time.Second)

			log.Printf("message %d done after %d seconds of wait!", v, dotCount)

			id.Add(1)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
