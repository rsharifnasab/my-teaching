package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Start listening on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	i := 0

	// Read data from the connection
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by client.")
			return
		}

		// Log the message from the client
		fmt.Printf("Received from client: %s", message)
		i += 1

		// Echo the message back to the client
		_, err = conn.Write([]byte("Echo: " + message + fmt.Sprintf("(%d)", i)))
		if err != nil {
			fmt.Println("Error sending message back to client:", err)
			return
		}
	}
}
