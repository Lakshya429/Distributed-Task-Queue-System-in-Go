package queues

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	ctx    *context.Context
	cancel *context.CancelFunc
	q      *amqp.Queue
}

var ConsumerBroker *Consumer

func Connection() error {
	conn, err := amqp.Dial("amqp://Lakshya:Lakshya123@localhost:5672/")
	if err != nil {
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	ConsumerBroker = &Consumer{
		ctx:    &ctx,
		cancel: &cancel,
		q:      &q,
	}

	defer cancel()

	return nil
}
func GetConsumer() *Consumer {
	return ConsumerBroker
}
