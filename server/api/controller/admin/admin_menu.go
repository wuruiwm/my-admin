package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func AdminMenuList(c *gin.Context) {
	response.Success(c, "success", logic.AdminMenuList())
}

func AdminMenuCreate(c *gin.Context) {
	param := &request.AdminMenuCreate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminMenuCreate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminMenuUpdate(c *gin.Context) {
	param := &request.AdminMenuUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminMenuUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminMenuDelete(c *gin.Context) {
	param := &request.AdminMenuDelete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminMenuDelete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminMenuSort(c *gin.Context) {
	param := &request.AdminMenuSort{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminMenuSort(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}
