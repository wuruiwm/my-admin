package main

import (
	"app/initialize"
)

func main() {
	initialize.Config()    //初始化配置
	initialize.Logger()    //初始化logger
	initialize.Db()        //初始化mysql连接池
	initialize.Redis()     //初始化redis连接池
	initialize.Validator() //初始化验证器翻译
	initialize.Task()      //初始化常驻任务
	initialize.CronTab()   //初始化定时任务
	initialize.Server()    //初始化httpServer
}
