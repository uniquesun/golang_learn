package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("启动服务端 ...")
	// network："tcp", "tcp4", "tcp6", "unix" or "unixpacket"
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("监听端口错误", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		// 第一次进入
		conn.Write([]byte("hello 欢迎进入聊天室"))
		// 回复客户
		go doServerWrite(conn, inputReader)
		// 接受客户消息
		go doServerRead(conn)
	}
}

func doServerWrite(conn net.Conn, inputReader *bufio.Reader) {
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误", err.Error())
			break
		}
		input = strings.TrimSpace(input) // 去除输入两端的空格
		if input == "exit" {
			fmt.Println("正在退出...")
			break
		}
		if len(input) > 0 {
			_, err = conn.Write([]byte(input))
			if err != nil {
				fmt.Println("发送数据错误", err.Error())
				break
			}
		}
	}
}

func doServerRead(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("读取客户数据错误", err1.Error())
			return
		}
		fmt.Printf("客户说：%v", string(buf[:n]))
	}

}
