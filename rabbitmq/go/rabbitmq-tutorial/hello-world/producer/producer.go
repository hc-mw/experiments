package main

import (
	"context"
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

	q := connObj.DeclareQueue("hello", false)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	connObj.PublishOnQueue(ctx, q.Name, "Hello, World!", "", false)
}
