package request

import "errors"

type PasswordList struct {
	PageBase
}

type PasswordCreate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
}

type PasswordUpdate struct {
	RequiredId
	PasswordCreate
}

type YoutubeRetry struct {
	RequiredId
}

type YoutubeDelete struct {
	RequiredId
}

func (param *PasswordCreate) Check() error {
	if param.Name == "" {
		return errors.New("请输入名称")
	}
	if param.Username == "" {
		return errors.New("请输入用户名")
	}
	if param.Password == "" {
		return errors.New("请输入密码")
	}
	return nil
}
