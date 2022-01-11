package main

import (
	"fmt"
	"net"
)


type Server struct {
	Ip string
	Port int
}

//創建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip: ip,
		Port: port,
	}

	return server
}


func(this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接建立成功")
}

//启动服务器的接口
func (this *Server) Start() {
	//socket Listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	//cloce listen socket
	defer listener.Close()

	for{
		//accpet
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listener accept err:",err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}


	

}