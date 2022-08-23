package admin

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/logic"
	"app/util"
	"github.com/gin-gonic/gin"
)

//AdminUserLogin 用户登录
func AdminUserLogin(c *gin.Context) {
	param := &request.AdminUserLogin{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if param.Username == "" || param.Password == "" {
		response.Error(c, "用户名或密码不能为空")
		return
	}

	//校验用户名密码是否正确 并登陆
	if result, err := logic.AdminUserLogin(param); err == nil {
		response.Success(c, "登陆成功", result)
	} else {
		response.Error(c, err.Error())
	}
}

//AdminUserMenu 获取用户的菜单
func AdminUserMenu(c *gin.Context) {
	userId := c.GetString("user_id")
	if result, err := logic.AdminUserMenu(userId); err == nil {
		response.Success(c, "success", result)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserList(c *gin.Context) {
	param := &request.AdminUserList{}
	if err := c.ShouldBindQuery(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if result, err := logic.AdminUserList(param); err == nil {
		response.Success(c, "success", result)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserCreate(c *gin.Context) {
	param := &request.AdminUserCreate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminUserCreate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserUpdate(c *gin.Context) {
	param := &request.AdminUserUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminUserUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserDelete(c *gin.Context) {
	param := &request.AdminUserDelete{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminUserDelete(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserRoleUpdate(c *gin.Context) {
	param := &request.AdminUserRoleUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminUserRoleUpdate(param); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserRole(c *gin.Context) {
	response.Success(c, "success", logic.AdminUserRole())
}

func AdminUserPasswordUpdate(c *gin.Context) {
	param := &request.AdminUserPasswordUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	if err := logic.AdminUserPasswordUpdate(param.Id, param.Password, global.Db); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserInfo(c *gin.Context) {
	userId := c.GetString("user_id")
	if result, err := logic.AdminUserInfo(userId); err == nil {
		response.Success(c, "success", result)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}

func AdminUserInfoUpdate(c *gin.Context) {
	param := &request.AdminUserInfoUpdate{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error(c, util.ValidatorError(err))
		return
	}

	userId := c.GetString("user_id")
	if err := logic.AdminUserInfoUpdate(param, userId); err == nil {
		response.Success(c, "success", nil)
	} else {
		c.Set("error", err.Error())
		response.Error(c, err.Error())
	}
}
