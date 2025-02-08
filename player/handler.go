package player

import (
	"fmt"
	"golang-gameserver/chat"
)

type Handler func(interface{})

func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFriend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg) // 这是一个接口类型，需要断言
	// 处理聊天消息
	fmt.Println(chatMsg)
	// todo 收到消息 然后转发给客户端（当你的好友给你发送消息时）
}
