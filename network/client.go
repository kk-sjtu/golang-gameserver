package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// 你作为客户端，有一个网关服务器
// 你有一个大厅的服务器。客户端想你请求，gateway充当客户端

// 这里需要一个类似客户端的身份，也就是client

type Client struct {
	Address string
	packer  NormalPacker
	// 客户端可以监听多个，所以不在这里放conn
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer: NormalPacker{
			Order: binary.BigEndian,
		},
	}
}
func (c *Client) Run() {
	// 连接服务器
	conn, err := net.Dial("tcp6", c.Address)
	if err != nil {
		fmt.Println("failed to connect", err)
		return
	}

	go c.Write(conn)
	go c.Read(conn)
	select {}
}
func (c Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			if conn == nil {
				fmt.Println("connection is nil, stopping write")
				return
			}
			fmt.Println("这是一个分割线")
			c.send(conn, &Message{
				Id:   666,
				Data: []byte("hello,lilithGame"),
			})

		}

	}
}

func (c *Client) send(conn net.Conn, message *Message) {
	if conn == nil {
		fmt.Println("connection is nil")
		return
	}
	err := conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		fmt.Println("connection is nil")
		return
	}
	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(bytes)
	if err != nil {

		fmt.Println("11111", err)
		return
	}
}

func (c *Client) Read(conn net.Conn) {
	if conn == nil {
		fmt.Println("connection is nil")
		return
	}
	//	err := conn.SetReadDeadline(time.Now().Add(time.Second))
	//	if err != nil {
	//		fmt.Println("connection is nil")
	//		return
	//	}
	//	for {
	//		message, err := c.packer.Unpack(conn)
	//		if _, ok := err.(net.Error); err != nil && ok {
	//			fmt.Println(err)
	//			continue
	//		}
	//		fmt.Println("client receive message:", string(message.Data))
	//	}
	//}

	for {
		err := conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // 增加超时时间
		if err != nil {
			fmt.Println(err)
			return
		}
		message, err := c.packer.Unpack(conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("111read timeout:", err)
				continue
			}
			if err.Error() == "1111connection reset by peer" {
				fmt.Println("1111connection reset by peer:", err)
				return
			}
			fmt.Println("1111unpack error:", err)
			return
		}
		fmt.Println("11111client receive message:", string(message.Data))
	}
}
