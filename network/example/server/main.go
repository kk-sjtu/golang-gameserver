package main

import "golang-gameserver/network"

func main() {
	server := network.NewServer(":8023", "tcp6")

	server.Run()
	select {}
}
