package router

import (
	"app/api/controller/api"
	"github.com/gin-gonic/gin"
)

// api路由
func apiRouter(r *gin.Engine) *gin.Engine {
	//路由组
	apiGroup := r.Group("/api")
	{
		uploadGroup := apiGroup.Group("/upload")
		{
			uploadGroup.POST("/image", api.UploadImage)
			uploadGroup.POST("/file", api.UploadFile)
		}
		//k8s登录
		apiGroup.POST("/k8s/login", api.K8sLogin)
		//ssl证书
		apiGroup.GET("/ssl", api.Ssl)
		//一言
		apiGroup.GET("/yiyan", api.Yiyan)
		//websocket
		apiGroup.GET("/websocket", api.Websocket)
		apiGroup.POST("/sendMsg", api.SendMsg)
	}
	return r
}
