package router

import (
	"app/api/controller/api"
	"github.com/gin-gonic/gin"
)

//api路由
func apiRouter(r *gin.Engine) *gin.Engine {
	//路由组
	apiGroup := r.Group("/api")
	{
		//上传
		apiGroup.POST("/upload/image", api.UploadImage)
		apiGroup.POST("/upload/file", api.UploadFile)
	}
	return r
}
