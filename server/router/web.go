package router

import (
	"app/api/controller/web"
	"github.com/gin-gonic/gin"
)

func webRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", web.Menu)
	return r
}
