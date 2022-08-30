package router

import (
	"app/api/controller/web"
	"github.com/gin-gonic/gin"
)

func webRouter(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("sdadasdsa"))
	})
	r.GET("/menu", web.Menu)
	return r
}
