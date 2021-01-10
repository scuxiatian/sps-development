package main

import (
	"sps-template-server/core"
	"sps-template-server/global"
	"sps-template-server/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main ()  {
	global.SdVp = core.Viper()
	global.SdLog = core.Zap()
	global.SdDB = initialize.GormMysql()
	initialize.MysqlTables(global.SdDB)

	db, _ := global.SdDB.DB()
	defer db.Close()

	core.RunServer()
}
