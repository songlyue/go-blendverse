package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func StartWebSocket() {
	http.HandleFunc("/acc", wsPage)

	// 添加处理程序

	http.ListenAndServe(":32475", nil)

}

func wsPage(req *http.Request, w http.ResponseWriter) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}
