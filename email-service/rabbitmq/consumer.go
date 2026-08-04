package rabbitmq

import (
	"encoding/json"
	"log"

	"example.com/mailer"
	"example.com/models"
)

func StartConsumer() {

	msgs, err := Channel.Consume(
		QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting for messages...")

	forever := make(chan bool)

	go func() {

		for msg := range msgs {

			var event models.ProjectCreatedEvent

			err := json.Unmarshal(msg.Body, &event)

			if err != nil {
				log.Println(err)
				continue
			}

			err = mailer.SendProjectConfirmation(event)

			if err != nil {
				log.Println(err)
			}
		}

	}()

	<-forever
}
