package request

import "errors"

type AdminMenuCreate struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	ParentId  string `json:"parent_id"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Title     string `json:"title"`
	Sort      uint   `json:"sort"`
	IsHidden  uint   `json:"is_hidden"`
}

type AdminMenuUpdate struct {
	RequiredId
	AdminMenuCreate
}

type AdminMenuDelete struct {
	RequiredId
}

type AdminMenuSort struct {
	RequiredId
	Type int `json:"type"`
}

func (param *AdminMenuCreate) Check() error {
	if param.Path == "" {
		return errors.New("请输入路由path")
	}
	if param.Name == "" {
		return errors.New("请输入路由name")
	}
	if param.Component == "" {
		return errors.New("请输入文件路径")
	}
	if param.IsHidden == 0 && param.Icon == "" {
		return errors.New("请选择图标")
	}
	if param.Title == "" {
		return errors.New("请输入菜单名称")
	}
	return nil
}
