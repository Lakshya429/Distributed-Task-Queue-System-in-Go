// config/rabbitmq.go
package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	QueueName   = "Wokers"
	ExchangeName = ""
	RoutingKey  = ""
)

// GetConnection returns a new RabbitMQ connection
func GetConnection() (*amqp.Connection, error) {
	// Update connection string as needed
	return amqp.Dial("amqp://guest:guest@localhost:5672/")
}

// GetChannel returns a new channel from a connection
func GetChannel() (*amqp.Channel, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	
	return conn.Channel()
}