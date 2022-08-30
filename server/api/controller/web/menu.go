package web

import (
	"app/api/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Menu(c *gin.Context) {
	if c.Request.Host == "menu.nikm.cn" {
		c.HTML(200, "menu.html", gin.H{
			"year": time.Now().Format("2006"),
		})
	} else {
		response.Error(c, fmt.Sprintf("api不存在 path:%s method:%s", c.Request.URL, c.Request.Method))
	}
}
