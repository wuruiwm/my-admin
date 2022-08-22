package initialize

import (
	"app/global"
	"github.com/spf13/viper"
)

//Config 初始化配置文件
func Config() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file error: " + err.Error())
	}
	readConfig()
}

func readConfig() {
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic("format config error:" + err.Error())
	}
}
