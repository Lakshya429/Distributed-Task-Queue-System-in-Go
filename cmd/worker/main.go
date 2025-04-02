package main 

import (
	"log"
	"github.com/Lakshya429/distributed-task-queue/internal/queues/consumer"
	"github.com/Lakshya429/distributed-task-queue/config"
)

func main() {
	ch , err  :=consumer.ConnectionConsumer()
	if err != nil {
		log.Fatal("Failed to get Conjumer Channel ", err)
	}
	msgs, err := ch.Consume(
		config.QueueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	  )

	  if err != nil {
		log.Fatal("Failed to get Conjumer Channel ", err)
	  }
	  forever := make(chan struct{})
	  go func() {
		for d := range msgs {
			log.Printf("Received a message : %s" , d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

