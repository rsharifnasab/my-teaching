package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publishMessage(ch *amqp.Channel, exchangeName, routingKey, body string) {
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func consumeMessages(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "test_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted when unused
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	queueName := "test_queue"
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	routingKey := "test_key"
	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		exchangeName, // exchange name
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to bind the queue to the exchange")

	body := "Hello RabbitMQ via Exchange!"
	publishMessage(ch, exchangeName, routingKey, body)

	// Simulate a delay before consuming
	time.Sleep(1 * time.Second)

	// Consume messages from the queue
	consumeMessages(ch, q.Name)
}
