package model

type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据
}
