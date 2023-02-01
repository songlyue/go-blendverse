package model

type Head struct {
	Seq      string    `json:"seq"`      // 消息的id
	Cmd      string    `json:"cmd"`      //cmd的动作
	Response *Response `json:"response"` //消息体
}

type Response struct {
	Code    uint32      `json:"code"`
	CodeMsg string      `json:"codeMsg"`
	Data    interface{} `json:"data"`
}

func NewResponse(code uint32, codeMsg string, data interface{}) *Response {
	return &Response{Code: code, CodeMsg: codeMsg, Data: data}
}

// 设置返回消息
func NewResponseHead(seq string, cmd string, code uint32, codeMsg string, data interface{}) *Head {
	response := NewResponse(code, codeMsg, data)

	return &Head{Seq: seq, Cmd: cmd, Response: response}
}
