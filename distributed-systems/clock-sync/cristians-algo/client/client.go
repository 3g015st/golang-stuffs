package main

import (
	"fmt"
	"net"
	"time"

	"github.com/araddon/dateparse"
)


func getServerTime() *time.Time {
	serverAddr := "127.0.0.1:8000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Failed to connect to the server:", err)
		return nil
	}

	defer conn.Close()

	requestTime := time.Now()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to receive data from the server:", err)
		return nil
	}

	serverTime, err := dateparse.ParseAny(string(buffer[:n]))

	if err != nil {
		fmt.Println("Failed to parse server time:", err)
		return nil
	}

	responseTime := time.Now()
	actualTime := time.Now()

	fmt.Println("Time returned by server:", serverTime)

	processDelayLatency := responseTime.Sub(requestTime).Seconds()
	elapsedTime := processDelayLatency / 2

	fmt.Println("Process Delay latency:", processDelayLatency, "seconds")
	fmt.Println("Actual clock time at client side:", actualTime)

	clientTime := serverTime.Add(time.Second * time.Duration(elapsedTime))

	error := actualTime.Sub(clientTime)
	fmt.Println("Synchronization error:", error.Seconds(), "seconds")

	return &clientTime
}

func main() {
	fmt.Println("Client Started")
	clientTime := getServerTime()
	fmt.Println("Synchronized process client time:", clientTime)
}
