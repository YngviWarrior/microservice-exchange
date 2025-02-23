package rabbitmq

import (
	"log"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type rabbitmq struct {
	Conn *amqp091.Connection
}

type RabbitMQInterface interface {
	Listen(string) <-chan amqp091.Delivery
}

func NewRabbitMQConnection() RabbitMQInterface {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão com RabbitMQ bem-sucedida!")

	return &rabbitmq{
		Conn: conn,
	}
}
