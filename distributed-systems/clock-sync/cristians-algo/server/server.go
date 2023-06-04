package main

import (
	"fmt"
	"net"
	"time"
)

const port = 8000

func initiateClockServer() {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Failed to start the server:", err)
		return
	}

	defer server.Close()

	fmt.Println("Socket is listening...")

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		fmt.Println("Server is connected to : ", connection.RemoteAddr())

		// Send server's time to client (UTC)
		connection.Write([]byte(time.Now().String()))

		connection.Close()
	}
}

func main() {
	fmt.Println("Server started")
	initiateClockServer()
}
