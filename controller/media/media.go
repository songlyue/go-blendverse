package media

import (
	"blendverse/model/common/response"
	"github.com/gin-gonic/gin"
)

func Start(c *gin.Context) {
	response.Ok(c)
}
