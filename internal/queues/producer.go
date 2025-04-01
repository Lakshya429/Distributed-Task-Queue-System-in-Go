package queue

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Producer represents a RabbitMQ message producer
type Producer struct {
	channel *amqp.Channel
}

// NewProducer initializes a new Producer
func NewProducer(rabbitMQURL string) (*Producer, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &Producer{channel: ch}, nil
}

// Publish sends a message to the specified queue
func (p *Producer) Publish(queueName string, message []byte) error {
	err := p.channel.Publish(
		"",         // Exchange (default)
		queueName,  // Routing key (queue name)
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Printf("Message published to queue: %s", queueName)
	return nil
}

// Close closes the producer's channel
func (p *Producer) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
}
