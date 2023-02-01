package websocket

import (
	"blendverse/global"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	clientManager = NewClientManager()
)

func StartWebSocket() {
	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go clientManager.start()
	http.ListenAndServe(":32475", nil)

}

func wsPage(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		global.GVA_LOG.Info(fmt.Sprintf("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"]))

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
	global.GVA_LOG.Info("客户端：%v 连接进来 ip：%s")
}
