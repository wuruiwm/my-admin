package initialize

import (
	"app/api/middleware"
	"app/global"
	"app/router"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

//Server 初始化gin的http服务
func Server(view embed.FS) {
	//是否开启debug
	if global.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//实例化gin
	r := gin.New()
	//Logger中间件
	r.Use(middleware.Logger)
	//Recovery中间件
	r.Use(middleware.Recover)
	//上传文件目录
	r.Static("/upload", "./resource/upload")
	t, err := template.ParseFS(view, "api/view/*.html")
	if err != nil {
		panic("view error:" + err.Error())
	}
	//html页面
	r.SetHTMLTemplate(t)
	//路由设置
	r = router.Router(r)
	//启动服务
	_ = r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.ServerPort))
}
