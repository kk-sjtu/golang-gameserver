package manager

import "golang-gameserver/player"

// playerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player
	addPCh  chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.UId] = p
	go p.Run()

}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}
