package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	//连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("conn=", conn)

	//获取命令行输入的信息
	reader := bufio.NewReader(os.Stdin)

	for {
		//获取客户端的单行信息
		line, err := reader.ReadString('\n')
		if err != nil {
			println(err)
		}
		//退出客户端
		if line == "exit\n" {
			println("exit succ")
			break
		}

		//发送数据到服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			println(err)
		}
		fmt.Println("发送成功", n)
	}



}