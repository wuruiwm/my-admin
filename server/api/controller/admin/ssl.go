package admin

import (
	"app/api/response"
	"app/logic"
	"github.com/gin-gonic/gin"
)

func Ssl(c *gin.Context) {
	if result, err := logic.SslList(); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}
