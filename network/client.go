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
		fmt.Println("client.go run is error", err)
		return
	}

	go c.Read(conn)
	go c.Write(conn)
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
			fmt.Println("你好了没")

		}
		fmt.Println("剑与远征")
	}

}

func (c *Client) send(conn net.Conn, message *Message) {

	err := conn.SetWriteDeadline(time.Now().Add(time.Second))

	if err != nil {
		fmt.Println("connection is nil")
		return
	}
	fmt.Println("弟弟")
	bytes, err := c.packer.Pack(message)
	fmt.Println("消息打包了吗")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("吃饭了吗")
	_, err = conn.Write(bytes)
	fmt.Println("这么持久")
	if err != nil {

		fmt.Println("2222", err)
	}
	fmt.Println("怎么回事")
}

func (c *Client) Read(conn net.Conn) {

	for {
		message, err := c.packer.Unpack(conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("resp message:", string(message.Data))
	}
}
