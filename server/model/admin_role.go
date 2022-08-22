package model

//AdminRole 后台角色表
type AdminRole struct {
	Base
	Name string `gorm:"column:name;type:char(20);comment:角色名称;NOT NULL" json:"name"`
}

func (m *AdminRole) TableName() string {
	return "admin_role"
}
