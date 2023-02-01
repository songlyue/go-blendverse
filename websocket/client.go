package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// 用户连接
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

// 初始化
func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}

	return
}

// 读取客户端的数据
func (c *Client) read() {

	defer func() {
		if recover() != nil {
			fmt.Println("write stop", recover())
		}
	}()

	defer func() {
		fmt.Println("读取客户端数据 关闭send", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取客户端数据错误", c.Addr, err)
			return
		}
		// 处理程序
		ProcessData(c, message)
	}

}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", r)
		}
	}()

	defer func() {
		// 注销client
		c.Socket.Close()
		fmt.Println("client发送数据", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				//发送数据失败
				fmt.Println("发送数据失败")
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}

}

// SendMsg 发送信息
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Send msg err")
		}
	}()
	c.Send <- msg
}
