## server.go

```go
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
```

## client.go

```go
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

```

## utils.go

```go
package main

import (
	"bytes"
	"encoding/binary"
	"net"
)

var (
	BufLen = 8
)

func SendMsg(conn net.Conn, message string) (n int, err error) {
	//获取消息的长度, 转化为[]byte类型
	var msgLen = make([]byte, 8)
	binary.BigEndian.PutUint32(msgLen, uint32(len(message)))

	//拼接消息长度和消息
	writeMsg := ByteAppend(msgLen[:BufLen], []byte(message))

	//向通道写入数据
	n, err = conn.Write(writeMsg)
	return
}

func ReadMsg(conn net.Conn) string {
	var buf [8096]byte
	_, _ = conn.Read(buf[:])
	msgLen := int(binary.BigEndian.Uint32(buf[:BufLen]))
	return string(buf[BufLen : msgLen+BufLen])
}

func ByteAppend(bts ...[]byte) []byte {
	var buffer bytes.Buffer
	for _, bt := range bts {
		buffer.Write(bt)
	}
	return buffer.Bytes()
}

```

