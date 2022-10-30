package client

import (
	"fmt"
	"ims/src/imsutils"
	"io"
	"net"
	"os"
)

type Client struct {
	Ip   string
	Port string
	Name string
	conn net.Conn
}

func NewClient(ip string, port string) *Client {
	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	//返回对象
	return &Client{
		Ip:   ip,
		Port: port,
		conn: conn,
	}
}
func (c *Client) run() {
	var msg string
	for {
		fmt.Print("请输入消息:")
		fmt.Scanln(&msg)
		sendMsg := c.Name + ":" + msg
		_, err := c.conn.Write([]byte(sendMsg))
		if err != nil {
			fmt.Println("conn Write err:", err)
		}
	}
}

// 处理server回应的消息， 直接显示到标准输出即可
func (c *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy到stdout标准输出上, 永久阻塞监听
	io.Copy(os.Stdout, c.conn)
}

// 服务器返回的消息
func (c *Client) GetRtnMsg() {
	buf := make([]byte, 4096)
	n, err := c.conn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Conn Read err:", err)
		return
	}
	str := string(buf[:n-1])
	msg_type, msg := imsutils.DealMsg(str)
	if msg_type != 0 {
		switch msg_type {
		case 1:
			//客户端用户名
			c.Name = msg
		case 2:
			//其他消息
			fmt.Println(msg)
		}
	}
}

func (c *Client) Start() {
	//处理服务器返回的消息
	c.GetRtnMsg()
	if len(c.Name) != 0 {
		fmt.Println("服务器连接成功，你好，" + c.Name + "！")
		//启动消息发送
		c.run()
	}
}
