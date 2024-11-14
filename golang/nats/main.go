package main

import (
	"fmt"
	"log"

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
	_ = message

	_, err = nc.Subscribe(
		subject,
		func(m *nats.Msg) {
			fmt.Printf("message subject: %s\n", m.Subject)
			fmt.Printf("Received message: %s\n", string(m.Data))
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
