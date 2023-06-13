package api

import (
	"app/api/response"
	"app/logic"
	"github.com/gin-gonic/gin"
)

func Pay(c *gin.Context) {
	if result, err := logic.Pay(); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}
