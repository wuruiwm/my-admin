package middleware

import (
	"app/api/response"
	"app/logic"
	"github.com/gin-gonic/gin"
)

func AdminUserAuth(c *gin.Context) {
	//获取token
	token := c.GetHeader("X-Admin-Token")
	//导出excel表格等接口 将token带到query参数上
	if token == "" {
		token = c.DefaultQuery("x_admin_token", "")
	}
	if adminUserLogin, code, err := logic.AdminUserAuth(token, c.Request.URL.Path, c.Request.Method); err != nil {
		c.Abort()
		c.Set("error", err.Error())
		response.Return(c, code, err.Error(), nil)
	} else {
		//将用户id存起来 以供上下文使用
		c.Set("user_id", adminUserLogin.Id)
		c.Next()
	}
}
