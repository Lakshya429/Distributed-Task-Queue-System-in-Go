package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Consumer represents a RabbitMQ message consumer
type Consumer struct {
	channel *amqp.Channel
}

// NewConsumer initializes a new Consumer
func NewConsumer(rabbitMQURL string) (*Consumer, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &Consumer{channel: ch}, nil
}

// Consume starts consuming messages from the specified queue
func (c *Consumer) Consume(queueName string, handler func([]byte)) error {
	msgs, err := c.channel.Consume(
		queueName,
		"",
		true,  // Auto-acknowledge
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()

	log.Printf("Consumer started for queue: %s", queueName)
	return nil
}

// Close closes the consumer's channel
func (c *Consumer) Close() {
	if c.channel != nil {
		c.channel.Close()
	}
}
