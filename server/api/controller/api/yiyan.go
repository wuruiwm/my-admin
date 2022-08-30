package api

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func Yiyan(c *gin.Context) {
	param := &request.Yiyan{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.Yiyan(param); err == nil {
		val, ok := result.(string)
		if param.ResponseType == "text" || param.ResponseType == "js" && ok {
			c.Writer.Write([]byte(val))
		} else {
			response.Success(c, "success", result)
		}
	} else {
		response.Error(c, err.Error())
	}
}
