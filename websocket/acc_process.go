package websocket

import (
	"blendverse/websocket/model"
	"encoding/json"
	"fmt"
	"sync"
)

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutes sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutes.Lock()
	defer handlersRWMutes.Unlock()
	handlers[key] = value
	return
}

// 获得方法
func getHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutes.Lock()
	defer handlersRWMutes.Unlock()
	value, ok = handlers[key]
	return
}

// ProcessData 处理数据
func ProcessData(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据stop", r)
		}
	}()
	fmt.Println("处理数据", client.Addr, string(message))

	request := &model.Request{}

	if err := json.Unmarshal(message, request); err != nil {
		fmt.Println("处理数据json 失败", err)
		client.SendMsg([]byte("处理json失败"))
		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		fmt.Println("处理json失败", err)
		client.SendMsg([]byte("处理json失败"))

		return
	}

	cmd := request.Cmd
	seq := request.Seq

	var (
		code uint32
		msg  string
		data interface{}
	)

	fmt.Println("acc_request", cmd, client.Addr)

	if handlerFunc, ok := getHandlers(cmd); ok {
		code, msg, data = handlerFunc(client, seq, requestData)
		msg = "ok"
	} else {
		fmt.Println("没有这个路由")
		msg = "err"
	}
	responseHead := model.NewResponseHead(seq, cmd, code, msg, data)
	headByte, err := json.Marshal(responseHead)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(headByte)

	fmt.Println("acc_response send", client.Addr, client.AppId, client.UserId, "cmd", cmd, "code", code)

	return
}
