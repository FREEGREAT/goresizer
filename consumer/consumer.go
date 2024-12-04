package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"goresizer.com/m/internal/utils"
)

const queueName = "Gompress01queue"

type MessageData struct {
	Message string  `json:"message"`
	Value   float64 `json:"value"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queueName, // ім'я черги
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Помилка оголошення черги: %v", err)
		return
	}

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Printf("Помилка отримання повідомлень: %v", err)
		return
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {

			var data MessageData
			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				log.Printf("Помилка десеріалізації повідомлення: %v", err)
				continue
			}

			log.Printf("Отримано повідомлення: %s з значенням: %f",
				data.Message,
				data.Value,
			)

			err = utils.Compress(data.Message, data.Value)
			if err != nil {
				log.Printf("Помилка обробки зображення: %v", err)
				continue
			}
		}
	}()
	log.Printf("Очікування повідомлень. Для виходу натисніть CTRL+C")
	<-forever
	return
}
