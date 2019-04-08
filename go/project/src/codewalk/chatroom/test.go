package main

import (
	"fmt"
	"net"
)


func listen(network string, address string)(err error) {
	//建立服务器进行端口监听
	listener, err := net.Listen(network, address)
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

		//客户端连接上来后, 读取客户端发送的消息
		var clientMess string
		for {
			//循环读取客户端提交的数据, 每次读取1024字节
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)

			//当n为0时, 代表读取完毕, 退出循环
			if n <= 0 {
				break
			}

			if err != nil {
				break
			}

			clientMess += string(buffer[:n])
		}

		//业务处理
		err = accept(conn, clientMess)
	}

	return
}


//客户端发送消息过来, 进行业务处理
func accept(conn net.Conn, message string)(err error) {
	fmt.Println("收到客户端", conn.RemoteAddr(),"发送的消息", message)
	return
}


func main()  {
	//监听端口
	err := listen("tcp", "0.0.0.0:8812")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
}
