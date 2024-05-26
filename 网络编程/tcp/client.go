package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("拨号失败", err.Error())
		return
	}
	defer conn.Close()

	buf := make([]byte, 512)
	inputReader := bufio.NewReader(os.Stdin)
	var wg sync.WaitGroup // 等待组，用于等待goroutine完成

	// 读取服务端消息的goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			n, err1 := conn.Read(buf)
			if err1 != nil {
				fmt.Println("读取服务数据错误", err1.Error())
				return
			}
			fmt.Printf(string(buf[:n]))
		}
	}()

	// 发送消息给服务端并监听退出命令
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

	// 等待读取消息的goroutine完成
	wg.Wait()
	fmt.Println("程序已退出")
}
