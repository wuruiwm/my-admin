package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode = 0 //通用成功 状态码
	ErrorCode   = 1 //通用错误 状态码

	LoginExpiredCode = -1 //用户鉴权失败
)

// Response 通用响应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 通用成功返回
func Success(c *gin.Context, msg string, data interface{}) {
	Return(c, SuccessCode, msg, data)
}

// Error 通用失败返回
func Error(c *gin.Context, msg string) {
	Return(c, ErrorCode, msg, nil)
}

// Return 通用返回
func Return(c *gin.Context, code int, msg string, data interface{}) {
	r := &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	Json(c, r)
}

// Json 通用json返回
func Json(c *gin.Context, r *Response) {
	c.JSON(http.StatusOK, r)
}

type Page struct {
	List     interface{} `json:"list"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
}

type Upload struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}
