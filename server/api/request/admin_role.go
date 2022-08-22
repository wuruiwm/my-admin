package request

import "errors"

type AdminRoleList struct {
	PageBase
}

type AdminRoleCreate struct {
	Name string `json:"name"`
}

type AdminRoleUpdate struct {
	RequiredId
	AdminRoleCreate
}

type AdminRoleDelete struct {
	RequiredId
}

type AdminRoleAuth struct {
	RequiredId
}

type AdminRoleAuthUpdate struct {
	RequiredId
	MenuId []string `json:"menu_id"`
	ApiId  []string `json:"api_id"`
}

func (param *AdminRoleCreate) Check() error {
	if param.Name == "" {
		return errors.New("请输入角色名称")
	}
	return nil
}
