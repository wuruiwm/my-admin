package api

import (
	"app/global"
	"github.com/gin-gonic/gin"
)

func DouyinCookie(c *gin.Context) {
	c.Writer.Write([]byte(global.Config.AdminConfig.Douyin.Cookie))
}
