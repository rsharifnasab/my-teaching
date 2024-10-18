package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server on localhost and port 8080
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Read input from the user and send it to the server
	for {
		fmt.Print("Enter message: ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		// Send the message to the server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Receive the response from the server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving response:", err)
			return
		}

		// Print the response from the server
		fmt.Print("Server response: " + response)
	}
}
