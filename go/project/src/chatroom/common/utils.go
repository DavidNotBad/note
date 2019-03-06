package common

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func ReadPkg(conn net.Conn)(mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Println("读到的长度为buf=", buf[:4])

	var pkgLen = binary.BigEndian.Uint32(buf[:4])
	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//把buf反序列化为Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func WritePkg(conn net.Conn, data []byte)(err error)  {
	//发送通讯内容的长度给对方， 用来验证是否丢包
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(len(data)))
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("发送消息的长度", len(data), "ok")

	//data就是可以发送的消息
	n, err = conn.Write(data)
	if n != int(len(data)) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("发送消息", string(data), "ok")
	return
}

