package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(broker string) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		return nil, err
	}
	return &Producer{producer: producer}, nil
}

func (p *Producer) Publish(topic string, key string, message string) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: -1},
		Key:            []byte(key),
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) Close() {
	p.producer.Close()
}
