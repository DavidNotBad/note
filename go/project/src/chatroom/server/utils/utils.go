package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}



func (this *Transfer) ReadPkg()(mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Println("读到的长度为buf=", this.Buf[:4])

	var pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//把buf反序列化为Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte)(err error)  {
	//发送通讯内容的长度给对方， 用来验证是否丢包
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], uint32(len(data)))
	n, err := this.Conn.Write(this.Buf[:4])
	fmt.Println("here..")
	fmt.Println(n)
	fmt.Println(string(this.Buf[:]))
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("发送消息的长度", len(data), "ok")

	//data就是可以发送的消息
	n, err = this.Conn.Write(data)
	if n != int(len(data)) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("发送消息", string(data), "ok")
	return
}




