package player

import "golang-gameserver/function"

type Player struct {
	UId        uint64
	FriendList []uint64 // 朋友
	chChat     chan chat.Msg
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: nil,
	}
	return p
}

func (p *Player) AddFriend(fId uint64) {
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

func (p *Player) DelFriend(fId uint64) {
	p.FriendList = function.DelEleInSlice(fld, p.FriendList)

}

func (p *Player) Run() {
	for {
		select {
		case chatMsg := <-p.chChat:
			// 处理聊天消息
			p.ResolveChatMsg(chatMsg)
		}
	}

}

func (p *Player) ResolveChatMsg(chatMsg chat.Msg) {

}
