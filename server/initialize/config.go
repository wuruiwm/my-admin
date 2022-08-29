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
		panic("config file read error: " + err.Error())
	}
	if err = viper.Unmarshal(&global.Config); err != nil {
		panic("config unmarshal error:" + err.Error())
	}
}
