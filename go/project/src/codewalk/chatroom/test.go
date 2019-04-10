package main

import (
	"fmt"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}

	//延时关闭服务器
	defer func() {
		err = listener.Close()
	}()

	for {
		//等待客户端连接上来, 这里会阻塞
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		//后面加上协程
		//读取包
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		fmt.Println("n=", n, "|content=", string(buffer[:n]), "|")

		//os.Exit(0)
		//业务处理
		//err = server.accept(conn, clientMess)
	}
}