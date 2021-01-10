package core

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"sps-template-server/global"
	"sps-template-server/initialize"
	"time"
)

func RunServer()  {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.SdConfig.System.Addr)
	server := http.Server{
		Addr:              address,
		Handler:           Router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	time.Sleep(10 * time.Microsecond)
	global.SdLog.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Sps-Dev
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8889
`, address)
	global.SdLog.Error(server.ListenAndServe().Error())
}
