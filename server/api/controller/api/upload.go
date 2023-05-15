package api

import (
	"app/api/response"
	"app/logic"
	"github.com/gin-gonic/gin"
)

// UploadImage 图片压缩上传
func UploadImage(c *gin.Context) {
	if result, err := logic.UploadImage(c, "file"); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}

// UploadFile 文件上传
func UploadFile(c *gin.Context) {
	if result, err := logic.UploadFile(c, "file"); err == nil {
		response.Success(c, "success", result)
	} else {
		response.Error(c, err.Error())
	}
}
