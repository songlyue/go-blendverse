package middleware

import (
	"blendverse/model/common/response"
	"blendverse/service/system"
	"blendverse/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = system.CasbinService{}

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUser, _ := utils.GetClaims(c)
		obj := c.Request.URL.Path
		act := c.Request.Method
		sub := strconv.Itoa(int(waitUser.AuthorityId))
		//	判断策略
		e := casbinService.Casbin()
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
		}
		c.Next()
	}
}
