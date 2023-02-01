package websocket

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	Clients     map[*Client]bool
	ClientsLock sync.RWMutex
	Users       map[string]*Client
	UserLock    sync.RWMutex
	Register    chan *Client
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)
		}

	}
}

func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

	// client.Send <- []byte("连接成功")
}

func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()

	manager.Clients[client] = true
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:  make(map[*Client]bool, 1000),
		Users:    make(map[string]*Client, 1000),
		Register: make(chan *Client, 1000),
	}
}
