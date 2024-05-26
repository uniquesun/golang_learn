package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello UDP server!"))
	if err != nil {
		fmt.Println("Error sending to UDP:", err)
		os.Exit(1)
	}

	// 接收回应（如果需要）
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		os.Exit(1)
	}

	// 打印收到的回应
	fmt.Println("Received from server:", string(buffer[:n]))
}
