package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"errors"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"net"
)

func K8sLogin(param *request.AdminUserLogin) (*response.K8sLogin, error) {
	if err := param.Check(); err != nil {
		return nil, err
	}
	if param.Username != global.Config.AdminConfig.K8s.AdminUsername {
		return nil, errors.New("该用户不允许登录k8s")
	}
	_, err := AdminUserCheckPassword(param.Username, param.Password, "id", "password", "salt", "admin_role_id")
	if err != nil {
		return nil, err
	}
	return K8sCreateToken()
}

func K8sCreateToken() (*response.K8sLogin, error) {
	config := &goph.Config{
		User:    global.Config.AdminConfig.K8s.Username,
		Addr:    global.Config.AdminConfig.K8s.Host,
		Port:    uint(global.Config.AdminConfig.K8s.Port),
		Auth:    goph.Password(global.Config.AdminConfig.K8s.Password),
		Timeout: goph.DefaultTimeout,
		Callback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := goph.NewConn(config)
	if err != nil {
		return nil, errors.New("ssh连接失败 error:" + err.Error())
	}
	out, err := client.Run("kubectl -n kubernetes-dashboard create token admin-user --duration 2592000s")
	if err != nil {
		return nil, errors.New("ssh命令执行失败 error:" + err.Error())
	}
	return &response.K8sLogin{
		Token: string(out),
	}, nil
}
