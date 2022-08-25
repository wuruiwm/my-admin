package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func AdminRoleList(c *gin.Context) {
	param := &request.AdminRoleList{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.AdminRoleList(param); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}

func AdminRoleCreate(c *gin.Context) {
	param := &request.AdminRoleCreate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminRoleCreate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func AdminRoleUpdate(c *gin.Context) {
	param := &request.AdminRoleUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminRoleUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func AdminRoleDelete(c *gin.Context) {
	param := &request.AdminRoleDelete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminRoleDelete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func AdminRoleAuth(c *gin.Context) {
	param := &request.AdminRoleAuth{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}
	response.Success(c, "success", logic.AdminRoleAuth(param))
}

func AdminRoleAuthUpdate(c *gin.Context) {
	param := &request.AdminRoleAuthUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminRoleAuthUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}
