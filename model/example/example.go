package example

import "blendverse/global"

type Example struct {
	global.GVA_MODEL
	Sb          string `json:"sb" gorm:"comment:sb"`               // api路径
	Path        string `json:"path" gorm:"comment:api路径"`          // api路径
	Description string `json:"description" gorm:"comment:api中文描述"` // api中文描述
}
