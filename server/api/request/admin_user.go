package request

import (
	"errors"
	"unicode/utf8"
)

type AdminUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminUserRoleId struct {
	AdminRoleId string `form:"admin_role_id" json:"admin_role_id"`
}

type AdminUserList struct {
	PageBase
	AdminUserRoleId
}

type AdminUserCreate struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	AdminUserRoleId
}

type AdminUserUpdate struct {
	RequiredId
	AdminUserCreate
}

type AdminUserDelete struct {
	RequiredId
}

type AdminUserRoleUpdate struct {
	RequiredId
	AdminUserRoleId
}

type AdminUserPasswordUpdate struct {
	RequiredId
	Password string `json:"password"`
}

type AdminUserInfoUpdate struct {
	Nickname    string `json:"nickname"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (param *AdminUserLogin) Check() error {
	if param.Username == "" {
		return errors.New("请填写用户名")
	}
	if utf8.RuneCountInString(param.Username) < 3 {
		return errors.New("用户名不能少于三位")
	}
	if param.Password != "" && utf8.RuneCountInString(param.Password) < 6 {
		return errors.New("密码不能少于六位")
	}
	return nil
}

func (param *AdminUserCreate) Check() error {
	if param.Username == "" {
		return errors.New("请填写用户名")
	}
	if utf8.RuneCountInString(param.Username) < 3 {
		return errors.New("用户名不能少于三位")
	}
	if param.Nickname == "" {
		return errors.New("请填写昵称")
	}
	if param.Password != "" && utf8.RuneCountInString(param.Password) < 6 {
		return errors.New("密码不能少于六位")
	}
	if param.AdminRoleId == "" {
		return errors.New("请选择角色")
	}
	return nil
}
