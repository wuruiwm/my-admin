package request

import "errors"

type K8sLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (param *K8sLogin) Check() error {
	if param.Username == "" {
		return errors.New("请输入用户名")
	}
	if param.Password == "" {
		return errors.New("请输入密码")
	}
	return nil
}
