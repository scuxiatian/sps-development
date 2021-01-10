package model

import "sps-template-server/global"

type JwtBlacklist struct {
	global.SdModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
