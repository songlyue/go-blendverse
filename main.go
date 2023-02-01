package main

import (
	"blendverse/config"
	"blendverse/global"
	"blendverse/initialize"
	"blendverse/router"
	"blendverse/websocket"
	webRouter "blendverse/websocket/router"
	"fmt"
)

func main() {
	// 配置
	initialize.Viper()
	// 日志
	global.GVA_LOG, _ = initialize.InitLogger(
		global.GVA_CONFIG.Log.LogFilePath,
		global.GVA_CONFIG.Log.LogInfoFileName,
		global.GVA_CONFIG.Log.LogWarnFileName,
		global.GVA_CONFIG.Log.LogFileExt,
	)
	// 数据库
	global.GVA_DB = initialize.InitGorm()
	initialize.RegisterTables(global.GVA_DB)
	config.NacosInit()
	// 参数校验翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	// 路由

	// webSocket 注册
	go websocket.StartWebSocket()
	webRouter.WebSocketInit()

	r := router.Router()
	global.GVA_LOG.Info("服务启动成功,%s")
	r.Run(fmt.Sprintf(":%s", global.GVA_CONFIG.Server.Port))
}
