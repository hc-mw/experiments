# Rabbit MQ Hello World
This is a simple hello-world project made while I was tinkering with RabbitMQ.

You can start consumer first, which will listen for messages to "hello" queue:
```bash
go run consumer/consumer.go
```

for producing messages, you can run `producer.go` for sending messages to `hello` queue.
The messages are ought to be provided via command line arg as shown below:
```bash
go run producer.go hey\ there
```
It will send the message and give output like below:
```bash
2025/08/21 18:53:30  [x] Sent "hey there"
```

That's it, mates.