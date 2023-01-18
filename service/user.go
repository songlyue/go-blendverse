package service

import (
	"blendverse/global"
	"blendverse/model/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRouter(g *gin.RouterGroup) {
	UserRouter := g.Group("user")
	{
		UserRouter.GET("list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"as": "as",
			})
		})
	}
}

var user system.User

func Login(param system.LoginParam) (*system.User, error) {
	rows := global.GVA_DB.Where(&system.User{UserName: param.UserName}).Find(&user)
	if rows.RowsAffected < 1 {
		return &user, nil
	}
	return &user, nil

}
