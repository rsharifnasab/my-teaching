package main

import (
	"fmt"
	"log"

	"fiber-bootstrap/bootstrap"
	"fiber-bootstrap/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()

	log.Fatal(
		app.Listen(
			fmt.Sprintf("%s:%s",
				env.GetEnv("APP_HOST", "localhost"),
				env.GetEnv("APP_PORT", "4000"),
			),
		),
	)
}
