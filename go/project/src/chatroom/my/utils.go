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
