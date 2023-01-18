package controller

import (
	"blendverse/model/common/response"
	"blendverse/model/system"
	"blendverse/service"
	"blendverse/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginParam system.LoginParam
	err := c.ShouldBind(&loginParam)
	if err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	login, err := service.Login(loginParam)
	if err != nil {
		response.FailWithDetailed(login, "异常", c)
	}
	response.OkWithData(login, c)
}
