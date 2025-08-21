package main

import (
	"context"
	"os"
	"rabbitmq-tut/utils"
	"time"
)

func main() {
	connObj := utils.NewRMQConn(utils.Credentials{
		User: "user",
		Pass: "password",
		Host: "localhost",
		Port: "5672",
	})
	defer connObj.CloseConnection()

	q := connObj.DeclareQueue("hello")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)

	connObj.PublishOnQueue(ctx, q.Name, body)
}
