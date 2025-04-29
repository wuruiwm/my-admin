package admin

import (
	"app/api/response"
	"app/util"
	"github.com/gin-gonic/gin"
)

func AliyunCdtFlow(c *gin.Context) {
	m := util.Map{}
	util.CacheGet("AliyunCdtFlow", &m)
	response.Success(c, "sucess", util.Map{
		"flow": m["flow"],
	})
}
