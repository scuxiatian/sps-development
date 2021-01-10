package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sps-template-server/global"
)

// 初始化viper 加载配置文件
func Viper() *viper.Viper  {
	config := "config.yaml"

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	// 监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 设置配置信息
		if err := v.Unmarshal(&global.SdConfig); err != nil {
			fmt.Println(err)
		}
	})

	// 设置配置信息
	if err := v.Unmarshal(&global.SdConfig); err != nil {
		fmt.Println(err)
	}
	return v
}
