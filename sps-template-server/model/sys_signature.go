package model

import "sps-template-server/global"

type SysSignature struct {
	global.SdModel
	Name		string		`json:"name" gorm:"comment:签章名"`
	Password	string		`json:"password" gorm:"comment:签章密码"`
	Url			string		`json:"url" gorm:"comment:签章图片地址"`
	Description	string		`json:"description" gorm:"comment:签章描述"`
	IsPublic	bool		`json:"isPublic" gorm:"comment:是否为公章"`
	Owner		SysUser		`json:"owner" gorm:"<-false" gorm:"comment:所属人(仅私章)"`
}
