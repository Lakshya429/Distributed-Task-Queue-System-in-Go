package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(broker, group string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	return &Consumer{consumer: c}, nil
}

func (c *Consumer) Subscribe(topic string) error {
	return c.consumer.Subscribe(topic, nil)
}

func (c *Consumer) StartProcessing(processFunc func(string)) {
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err == nil {
			processFunc(string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
