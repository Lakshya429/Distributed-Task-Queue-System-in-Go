package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQConfig holds the connection and channel
type RabbitMQConfig struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewRabbitMQConfig initializes a RabbitMQ connection and channel
func NewRabbitMQConfig(rabbitMQURL string) (*RabbitMQConfig, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQConfig{
		Connection: conn,
		Channel:    ch,
	}, nil
}

// DeclareQueue declares a queue with the given name
func (r *RabbitMQConfig) DeclareQueue(queueName string) error {
	_, err := r.Channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	return nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQConfig) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}
