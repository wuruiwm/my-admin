package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func Ssl(c *gin.Context) {
	param := &request.Ssl{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}
	if result, err := logic.Ssl(param); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}
