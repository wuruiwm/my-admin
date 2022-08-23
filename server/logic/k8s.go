package logic

import (
	"app/api/request"
	"app/api/response"
	"app/util"
	"errors"
	"fmt"
	"github.com/melbahja/goph"
	"strconv"
)

func K8sLogin(param *request.AdminUserLogin) (interface{}, error) {
	if err := param.Check(); err != nil {
		return nil, err
	}
	if param.Username != util.AdminConfig("k8s.admin_user") {
		return nil, errors.New("该用户不允许登录k8s")
	}
	_, err := AdminUserCheckPassword(param.Username, param.Password, "id", "password", "salt", "admin_role_id")
	if err != nil {
		return nil, err
	}
	return K8sCreateToken()
}

func K8sCreateToken() (*response.K8sLogin, error) {
	callback, err := goph.DefaultKnownHosts()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(util.AdminConfig("k8s.port"))
	if err != nil {
		return nil, errors.New("端口参数不正确 error:" + err.Error())
	}
	client, err := goph.NewConn(&goph.Config{
		User:     util.AdminConfig("k8s.user"),
		Addr:     util.AdminConfig("k8s.host"),
		Port:     uint(port),
		Auth:     goph.Password(util.AdminConfig("k8s.password")),
		Timeout:  goph.DefaultTimeout,
		Callback: callback,
	})
	if err != nil {
		return nil, errors.New("ssh连接失败 error:" + err.Error())
	}
	out, err := client.Run("kubectl -n kubernetes-dashboard create token admin-user --duration 2592000s")
	if err != nil {
		return nil, errors.New("ssh命令执行失败 error:" + err.Error())
	}
	fmt.Println(string(out))
	return &response.K8sLogin{
		Token: string(out),
	}, nil
}
