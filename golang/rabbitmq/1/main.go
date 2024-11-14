package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// failOnError is a helper function to check for errors and log them
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// publishMessage publishes a message to the specified queue
func publishMessage(ch *amqp.Channel, queueName, body string) {
	err := ch.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

// consumeMessages sets up a consumer to receive messages from the queue
func consumeMessages(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {

			time.Sleep(1 * time.Second)
			log.Printf("Received a message: %s", d.Body)
			// d.Ack(false)
			d.Nack(false, true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close() // TODO: check err

	queueName := "test_queue"

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			//	"queue-type":  "classic",
			"message-ttl": 10_000_000,
		},
	)
	println(q.Name)
	failOnError(err, "Failed to declare a queue")

	body := "Hello RabbitMQ!"

	consumeMessages(ch, queueName)
	for range 10 {
		publishMessage(ch, queueName, body)
	}
	select {}
}
