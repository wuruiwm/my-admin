package admin

import (
	"app/api/request"
	"app/api/response"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

func YoutubeCreate(c *gin.Context) {
	param := &request.YoutubeCreate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.YoutubeCreate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func YoutubeList(c *gin.Context) {
	param := &request.YoutubeList{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.YoutubeList(param); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}

func YoutubeRetry(c *gin.Context) {
	param := &request.YoutubeRetry{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.YoutubeRetry(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}

func YoutubeDelete(c *gin.Context) {
	param := &request.YoutubeDelete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.YoutubeDelete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		response.Error(c, err.Error())
	}
}
