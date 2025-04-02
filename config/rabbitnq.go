package config

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	QueueName   = "Workers" // Fixed typo from "Wokers"
	ExchangeName = ""
	RoutingKey  = ""
)

// GetChannel returns a new RabbitMQ connection and channel
func GetChannel() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://Lakshya:Lakshya123@localhost:5672/")
	if err != nil {
		log.Printf("Failed to get Connection: %v", err)
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close() // Clean up connection if channel creation fails
		return nil, nil, err
	}

	return conn, ch, nil
}