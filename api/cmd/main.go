package main

import (
	"example.com/rabbitmq"
	"example.com/routes"
)

func main() {

	rabbitmq.Connect()
	router := routes.SetupRouter()

	router.Run(":8080")
}
