package consumer

import (
	"log"
	"errors"
	"github.com/Lakshya429/distributed-task-queue/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
var ch *amqp.Channel
var err error
func ConnectionConsumer() error {
	ch , err = config.GetChannel()

	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		config.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	return nil
}

func GetConsumer()(*amqp.Channel , error){
	if (ch == nil){
		return nil, errors.New("channel not initialized")
	}
	return ch , nil
}


