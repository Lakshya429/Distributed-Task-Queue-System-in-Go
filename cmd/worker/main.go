package main

import (
	"log"

	"github.com/Lakshya429/distributed-task-queue/internal/repository"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	repository.Setup()

	consumer, err := confluentkafka.NewConsumer(&confluentkafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "video-processor-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	consumer.Subscribe("video_events", nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v", err)
			continue
		}
		log.Printf("Received message: %s", string(msg.Value))

	}

}
