package api

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func K8sLogin(c *gin.Context) {
	param := &request.AdminUserLogin{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.K8sLogin(param); err == nil {
		response.Success(c, "success", result)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}
