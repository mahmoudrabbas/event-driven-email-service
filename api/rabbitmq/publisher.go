package rabbitmq

import (
	"encoding/json"

	"example.com/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(event shared.ProjectCreatedEvent) error {

	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	return Channel.Publish("", QueueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

}
