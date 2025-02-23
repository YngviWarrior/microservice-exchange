package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func (r *rabbitmq) Listen(queueName string) (msgs <-chan amqp091.Delivery) {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		queue.Name, // Nome da fila
		"",         // Routing Key (não usada em Fanout)
		"klines",   // Nome do Exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err = ch.Consume(
		queue.Name,
		"",
		true,  // Auto Ack: true (confirma automaticamente)
		false, // Exclusive: false
		false, // No Local: false
		false, // No Wait: false
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		fmt.Println("Recebido:", string(msg.Body))
	}
	return
}
