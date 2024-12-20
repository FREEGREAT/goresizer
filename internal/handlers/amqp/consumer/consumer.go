package consumer

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"goresizer.com/m/internal/config"
	service "goresizer.com/m/internal/service"
)

type MessageData struct {
	Message string  `json:"message"`
	Value   float64 `json:"value"`
}

func Consumer() {
	cfg := config.GetConfig().RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Помилка підключення до RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Помилка створення каналу: %v", err)
		return
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		cfg.QueueName, // ім'я черги
		true,          // durable
		false,         // autoDelete
		false,         // exclusive
		false,         // noWait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Помилка оголошення черги: %v", err)
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
		log.Fatalf("Помилка отримання повідомлень: %v", err)
		return
	}

	for msg := range msgs {
		var data MessageData
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			log.Printf("Помилка десеріалізації повідомлення: %v", err)
			continue
		}

		log.Printf("Отримано повідомлення: %s з значенням: %f", data.Message, data.Value)

		err = service.Compress(data.Message, data.Value)
		if err != nil {
			log.Printf("Помилка обробки зображення: %v", err)
		} else {
			log.Printf("Повідомлення успішно оброблено: %s", data.Message)
		}
	}
}
