package model

//AdminRoleMenu 后台角色与菜单关联表
type AdminRoleMenu struct {
	Base
	AdminRoleId string `gorm:"column:admin_role_id;type:char(36);comment:管理员id;NOT NULL" json:"admin_role_id"`
	AdminMenuId string `gorm:"column:admin_menu_id;type:char(36);comment:菜单id;NOT NULL" json:"admin_menu_id"`
}

func (m *AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}
