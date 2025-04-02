package consumer

import (
	"github.com/Lakshya429/distributed-task-queue/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Global variables for connection and channel
var conn *amqp.Connection
var ch *amqp.Channel

func ConnectionConsumer() ( *amqp.Channel , error) {
	var err error
	conn, ch, err = config.GetChannel()
	if err != nil {
		return ch , err
	}

	_, err = ch.QueueDeclare(
		config.QueueName, // "Workers"
		false,            // durable
		false,            // auto-delete
		false,            // exclusive
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return ch , err
	}
	return ch ,nil
}
// Close cleans up the connection and channel
func Close() {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}