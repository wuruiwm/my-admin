package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func M3u8Create(c *gin.Context) {
	param := &request.M3u8Create{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.M3u8Create(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func M3u8List(c *gin.Context) {
	param := &request.M3u8List{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.M3u8List(param); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}

func M3u8Retry(c *gin.Context) {
	param := &request.M3u8Retry{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.M3u8Retry(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func M3u8Delete(c *gin.Context) {
	param := &request.M3u8Delete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.M3u8Delete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}
