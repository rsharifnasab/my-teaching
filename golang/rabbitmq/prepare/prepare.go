package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// failOnError is a helper function to check for errors and log them
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	defer func() {
		err := ch.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	queueName := "test_queue"

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			//	"queue-type":  "classic",
			"x-message-ttl": 10_000,
		},
	)
	println(q.Name)
	failOnError(err, "Failed to declare a queue")
}
