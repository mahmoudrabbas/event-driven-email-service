package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var Channel *amqp.Channel

const QueueName = "project_created"

func Connect() {

	var err error

	Conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = Channel.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email Service Connected")
}
