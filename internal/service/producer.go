package service

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"goresizer.com/m/internal/config"
)

var cfg = config.GetConfig()

type MessageData struct {
	Message string  `json:"message"`
	Value   float64 `json:"value"`
}

func PublishMessage(message string, value float64) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		cfg.RabbitMQ.QueueName, // queue name
		true,                   // durable
		false,                  // autoDelete
		false,                  // exclusive
		false,                  // noWait
		nil,                    // arguments
	)
	if err != nil {
		log.Printf("Cannot create queue: %v", err)
		return err
	}

	data := MessageData{
		Message: message,
		Value:   value,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error data serialize: %v", err)
		return err
	}

	err = channel.Publish(
		"",                     // exchange
		cfg.RabbitMQ.QueueName, // routing key
		false,                  // mandatory
		false,                  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)
	if err != nil {
		log.Printf("message sending error: %v", err)
		return err
	}

	log.Printf("Message send: %s with values: %f", message, value)
	return nil
}
