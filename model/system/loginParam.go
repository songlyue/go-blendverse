package system

type LoginParam struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	PassWord string `form:"passWord" json:"passWord" binding:"required"`
}
