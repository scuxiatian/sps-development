package service

import (
	"sps-template-server/global"
	"sps-template-server/model"
)

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.SdDB.Create(&jwtList).Error
	return err
}
