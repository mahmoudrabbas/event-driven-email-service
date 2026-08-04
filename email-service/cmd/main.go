package main

import (
	"example.com/config"
	"example.com/rabbitmq"
)

func main() {

	config.Load()

	rabbitmq.Connect()

	rabbitmq.StartConsumer()
}
