package model

import (
	"sps-template-server/global"
)

type SysSignature struct {
	global.SdModel
	Name		string		`json:"name" gorm:"comment:签章名"`
	Password	string		`json:"password" gorm:"comment:签章密码"`
	Url			string		`json:"url" gorm:"comment:签章图片地址"`
	Description	string		`json:"description" gorm:"comment:签章描述"`
	IsPublic	bool		`json:"isPublic" gorm:"comment:是否为公章"`
	Owner		SysUser		`json:"owner" gorm:"foreignKey:UUID;references:OwnerId;comment:所属人(仅私章)"`
	OwnerId		string		`json:"ownerId" gorm:"comment:所属人ID"`
}

type SignatureRecord struct {
	global.SdModel
	SignatureUses	[]SignatureUse	`json:"signatures"`
}

type SignatureUse struct {
	global.SdModel
	SignatureRecordID	uint			`json:"signatureRecordID" gorm:"comment:签章记录ID"`
	SignatureBase		SysSignature	`json:"signature" gorm:"foreignKey:ID;references:SignatureId;comment:签章"`
	SignatureId 		uint			`json:"signatureId" gorm:"comment:签章ID"`
	X					uint			`json:"x" gorm:"default:0;comment:签章x轴偏移"`
	Y					uint			`json:"y" gorm:"default:0;comment:签章y轴偏移"`
	Operator			SysUser			`json:"operator" grom:"foreignKey:ID;references:OperatorID;comment:签章人"`
	OperatorID			uint			`json:"operatorId" grom:"comment:签章人ID"`
	Description			string			`json:"description" grom:"comment:说明"`
}
