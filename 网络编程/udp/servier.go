package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 监听本地端口
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	// 创建一个UDP连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", addr.String())

	// 缓冲区，用于接收数据
	buffer := make([]byte, 1024)

	for {
		// 读取数据
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// 打印收到的数据
		fmt.Printf("Received from %s: %s\n", addr.String(), string(buffer[:n]))

		// 发送回应（如果需要）
		_, err = conn.WriteToUDP([]byte("Message received"), addr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
		}
	}
}
