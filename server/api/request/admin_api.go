package request

import (
	"errors"
	"slices"
)

type AdminApiList struct {
	PageBase
	Group  string `form:"group"`
	Method string `form:"method"`
}

type AdminApiCreate struct {
	Name   string `json:"name"`
	Group  string `json:"group"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type AdminApiUpdate struct {
	RequiredId
	AdminApiCreate
}

type AdminApiDelete struct {
	RequiredId
}

var requestMethod = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
}

func (param *AdminApiCreate) Check() error {
	if param.Name == "" {
		return errors.New("请输入api名称")
	}
	if param.Group == "" {
		return errors.New("请输入分组")
	}
	if param.Path == "" {
		return errors.New("请输入接口路径")
	}
	if param.Method == "" {
		return errors.New("请选择请求方式")
	}
	if !slices.Contains(requestMethod, param.Method) {
		return errors.New("请选择正确的请求方式")
	}
	return nil
}
