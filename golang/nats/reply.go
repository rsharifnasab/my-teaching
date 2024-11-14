package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "service.request"

	go func() {
		nc.Subscribe(subject, func(m *nats.Msg) {
			fmt.Printf("Received request: %s\n", string(m.Data))

			// Can by dynamic
			response := "Here is your response!"
			err := nc.Publish(m.Reply, []byte(response))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Sent reply: %s\n", response)
		})

		// Keep the listener running
		select {}
	}()

	time.Sleep(1 * time.Second)

	msg, err := nc.Request(subject, []byte("Hello, can I get a response?"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received reply: %s\n", string(msg.Data))

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
}
