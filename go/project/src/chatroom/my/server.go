package main

import (
	"fmt"
	"net"
)

func main() {
	//监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Printf("net.Listen error: %s", err)
		return
	}
	//延时关闭
	defer func() {
		if err := listener.Close(); err != nil {
			fmt.Printf("listener.Close error: %s", err)
			return
		}
	}()

	//等待连接
	for {
		//阻塞
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener.Accept error: %s", err)
			continue
		}

		//连接成功后， 启动协程和客户端保持通讯
		go process(conn)
	}
}

func process(conn net.Conn) {
	//延时关闭
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("conn.Close error: %s", err)
			return
		}
	}()

	fmt.Println(ReadMsg(conn))
}
