package router

import "blendverse/websocket"

func WebSocketInit() {
	websocket.Register("ping", websocket.Ping)
}
