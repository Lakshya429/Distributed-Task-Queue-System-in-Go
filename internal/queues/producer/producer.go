package producer


import (
	"context"
	"encoding/json"
	"time"
	"github.com/Lakshya429/distributed-task-queue/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Lakshya429/distributed-task-queue/internal/models"
)

// Global variables for connection and channel
var conn *amqp.Connection
var ch *amqp.Channel

func ConnectionProducer() ( *amqp.Channel , error) {
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
func PublishMessage(videoRequest *models.VideoRequest) error {
	jsonData , err := json.Marshal(videoRequest)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
  defer cancel()
  
  err = ch.PublishWithContext(
	ctx,
	"",     // exchange
    "Workers", // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing {
	  ContentType: "application/json",
	  Body:        jsonData,
	})
	if err != nil {
		return err
	}
	return nil
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