package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Printf("等待%v输入。。。", conn.RemoteAddr().String())
		//等待客户端发送数据
		n, err := conn.Read(buf)
		if err == io.EOF {
			println("客户端已退出", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}


func main() {
	fmt.Println("服务器开始监听了。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		println(conn)

		go process(conn)
	}

}