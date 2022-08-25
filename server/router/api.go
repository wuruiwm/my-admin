package router

import (
	"app/api/controller/admin"
	"app/api/controller/api"
	"github.com/gin-gonic/gin"
)

//api路由
func apiRouter(r *gin.Engine) *gin.Engine {
	//路由组
	apiGroup := r.Group("/api")
	{
		//上传
		uploadGroup := apiGroup.Group("/upload")
		{
			uploadGroup.POST("/image", api.UploadImage)
			uploadGroup.POST("/file", api.UploadFile)
		}
		//k8s登录
		apiGroup.POST("/k8s/login", api.K8sLogin)
		//ssl证书
		apiGroup.GET("/ssl", admin.Ssl)
	}
	return r
}
