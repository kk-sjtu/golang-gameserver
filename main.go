package main

import "golang-gameserver/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
