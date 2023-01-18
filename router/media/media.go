package media

import (
	"blendverse/controller/media"
	"github.com/gin-gonic/gin"
)

func Router(group *gin.RouterGroup) {

	mediaGroup := group.Group("v2")
	mediaGroup.POST("/group_photo/join_photo_info", media.Start)
	mediaGroup.POST("/group_photo/add_photo_users", media.Start)
	mediaGroup.POST("/group_photo/delete_photo_users", media.Start)
	mediaGroup.POST("/group_photo/update_user_info", media.Start)

}
