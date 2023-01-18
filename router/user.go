package router

import (
	"blendverse/controller"
	"blendverse/router/media"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	group := r.Group("/api/")
	media.Router(group)
	group.POST("/login", controller.Login)
	return r
}
