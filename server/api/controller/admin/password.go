package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func PasswordList(c *gin.Context) {
	param := &request.PasswordList{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.PasswordList(param); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}

func PasswordCreate(c *gin.Context) {
	param := &request.PasswordCreate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.PasswordCreate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func PasswordUpdate(c *gin.Context) {
	param := &request.PasswordUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.PasswordUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func PasswordDelete(c *gin.Context) {
	param := &request.PasswordDelete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.PasswordDelete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}
