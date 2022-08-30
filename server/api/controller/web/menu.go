package web

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Menu(c *gin.Context) {
	c.HTML(200, "menu.html", gin.H{
		"year": time.Now().Format("2006"),
	})
}
