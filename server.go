package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port string
}

func NewServer(ip string, port string) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	fmt.Println("链接建立成功")
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	//close listen socket
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		go this.Handler(conn)
	}

}
