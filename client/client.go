package main

import (
	"encoding/json"
	"fmt"
	"golang-gameserver/network"
)

type Client struct {
	Cli             *network.Client
	inputHandlers   map[string]InputHandler   // 收发消息，从控制台输入消息。充当用户玩家行为
	messageHandlers map[string]MessageHandler // 处理服务端发过来的消息
	console         *ClientConsole            // 就是控制台，理解为一个类
	chInput         chan *Inputparam          // 处理控制台，跑数据
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[uint64]MessageHandler{},
		console:         NewClientConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}
func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				fmt.Printf("cmd:%s,param:%v  <<<\t \n", input.Command, input.Param)
				bytes, err := json.Marshal(input.Param)
				if err == nil {
					c.cli.ChMsg <- &network.Message{
						ID:   1,
						Data: bytes,
					}
				}
			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}
func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[packet.Msg.ID]; ok {
		handler(packet)
	}
}
