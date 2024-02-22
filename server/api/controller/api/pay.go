package api

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func Pay(c *gin.Context) {
	param := &request.Pay{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}
	if param.Password != global.Config.AdminConfig.Pay.Password {
		response.Error(c, "密码错误")
		return
	}
	if result, err := logic.Pay(); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}
