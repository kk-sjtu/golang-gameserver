package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn    net.Conn
	packer  *NormalPacker
	chWrite chan *Message
	//
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		chWrite: make(chan *Message, 1),
	}
}

func (s *Session) Run() { // 这是一个会话。例如现在访问一个网站
	// ，给你一个session维护起来。有一个生命周期。例如淘宝，关掉一页，登录状态一直也在
	// todo 客户端收发数据

	go s.Read()
	go s.Write()

}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		message, err := s.packer.Unpack(s.conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("server receive message:", string(message.Data))
		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("hi ,lilith"),
		}

	}

	// todo 接收的是字节流，需要处理
}

func (s *Session) Write() {
	// todo 发送的是字节流，需要处理
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case msg := <-s.chWrite:
			s.send(msg)

		}

	}

}

func (s *Session) send(message *Message) {
	bytes, err := s.packer.Pack(message)
	if err != nil {
		return
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}

}
