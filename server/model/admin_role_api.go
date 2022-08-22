package model

//AdminRoleApi 后台角色与api关联表
type AdminRoleApi struct {
	Base
	AdminRoleId string `gorm:"column:admin_role_id;type:char(36);comment:角色id;NOT NULL" json:"admin_role_id"`
	AdminApiId  string `gorm:"column:admin_api_id;type:char(36);comment:api id;NOT NULL" json:"admin_api_id"`
}

func (m *AdminRoleApi) TableName() string {
	return "admin_role_api"
}
