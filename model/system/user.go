package system

import "blendverse/global"

type User struct {
	global.GVA_MODEL
	UserName string `form:"userName" json:"userName" binding:"required"`
	PassWord string `form:"passWord" json:"passWord" binding:"required"`
}
