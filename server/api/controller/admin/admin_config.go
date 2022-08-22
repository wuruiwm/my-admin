package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func AdminConfigList(c *gin.Context) {
	param := &request.AdminConfigList{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}
	response.Success(c, "success", logic.AdminConfigList(param))
}

func AdminConfigUpdate(c *gin.Context) {
	param := &request.AdminConfigUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminConfigUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}
