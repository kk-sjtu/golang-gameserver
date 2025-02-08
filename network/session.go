package network

import (
	"net"
	"time"
)

type Session struct {
	conn net.Conn
	//
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}

func (s *Session) Run() { // 这是一个会话。例如现在访问一个网站
	// ，给你一个session维护起来。有一个生命周期。例如淘宝，关掉一页，登录状态一直也在
	// todo 客户端收发数据

}

func (s *Session) Read() {
	err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.println(err)
	}

	// todo 接收的是字节流，需要处理
}

func (s *Session) Write() {
	// todo 发送的是字节流，需要处理
}
