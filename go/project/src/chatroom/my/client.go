package main

import (
	"fmt"
	"net"
)

func main() {
	//连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Printf("net.Dial error: %s", err)
		return
	}
	//延时关闭
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("conn.Close error: %s", err)
			return
		}
	}()

	//发送消息
	_, err = SendMsg(conn, "hello")
	if err != nil {
		fmt.Printf("Send error: %s", err)
	}

	msg := ReadMsg(conn)
	fmt.Println(msg)
}
