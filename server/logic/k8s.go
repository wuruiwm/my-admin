package logic

import "app/api/request"

func K8sLogin(param *request.K8sLogin) {
	if err := param.Check(); err != nil {
		return err
	}
}
