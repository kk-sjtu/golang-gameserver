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
			fmt.Println("这是一个分割线,说明conn不为空")
			c.send(conn, &Message{
				Id:   666,
				Data: []byte("hello,lilith Game"),
			})

		}

	}
}

func (c *Client) send(conn net.Conn, message *Message) {

	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
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

		fmt.Println("2222", err)
	}
}

func (c *Client) Read(conn net.Conn) {

	for {
		message, err := c.packer.Unpack(conn)
		if _, ok := err.(net.Error); err != nil && ok {
			fmt.Println(err)
			continue
		}
		fmt.Println("resp message:", string(message.Data))
	}
}
