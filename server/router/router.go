package router

import (
	"github.com/gin-gonic/gin"
)

//Router 设置路由
func Router(r *gin.Engine) *gin.Engine {
	r = adminRouter(r)
	r = apiRouter(r)
	return r
}
