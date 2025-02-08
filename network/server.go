package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	address  string
	network  string
}

func NewServer(address, network string) *Server {
	return &Server{
		address:  address,
		network:  network,
		listener: nil,
	}
}

func (s *Server) Run() {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.address)
	if err != nil {
		fmt.Println("ResolveTCPAddr error:", err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		fmt.Println("ListenTCP error:", err)
		return
	}
	s.listener = tcpListener
	
}
