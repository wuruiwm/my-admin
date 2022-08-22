package middleware

import (
	"app/api/response"
	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.Set("panic", err)
			response.Error(c, "接口异常捕获")
		}
	}()
	c.Next()
}
