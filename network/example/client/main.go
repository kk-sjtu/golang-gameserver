package main

import "golang-gameserver/network"

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}
