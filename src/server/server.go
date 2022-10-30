package server

import (
	"fmt"
	"ims/src/imsutils"
	"net"
)

type Server struct {
	IP      string
	Port    string
	Message chan string
}

func NewServer(ip string, port string) *Server {
	return &Server{
		IP:      ip,
		Port:    port,
		Message: make(chan string),
	}
}
func (se *Server) ReadMsg(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, _ := conn.Read(buf)
		if n == 0 {
			return
		}
		msg := string(buf[:n])
		fmt.Println(msg)
	}
}

func (se *Server) HadelService(conn net.Conn) {
	//将用户名发送客户端
	cln := imsutils.GetRandstring(5)
	sendMsg := "newclient:" + cln
	_, err := conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
	}
	//监听本端口的所有消息，放入管道
	go se.ReadMsg(conn)
}
func (se *Server) Start() {
	fmt.Println("server start...")
	//监听端口，有消息直接广播
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", se.IP, se.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//方法结束关闭
	defer listener.Close()
	//监听服务端口
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
		}
		go se.HadelService(conn)
	}
}
