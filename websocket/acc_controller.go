package websocket

func Ping(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {
	code = 200
	data = "pong"
	return
}
