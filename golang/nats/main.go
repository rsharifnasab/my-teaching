package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL) // DefaultURL is nats://localhost:4222
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "updates"
	message := "Hello, NATS!"

	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	err = nc.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published message: %s\n", message)

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}
