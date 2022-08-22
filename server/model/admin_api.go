package model

//AdminApi 后台api表
type AdminApi struct {
	Base
	Name   string `gorm:"column:name;type:char(20);comment:接口名称;NOT NULL" json:"name"`
	Group  string `gorm:"column:group;type:char(20);comment:分组;NOT NULL" json:"group"`
	Path   string `gorm:"column:path;type:char(50);comment:接口路径;NOT NULL" json:"path"`
	Method string `gorm:"column:method;type:char(5);comment:请求方式;NOT NULL" json:"method"`
}

func (m *AdminApi) TableName() string {
	return "admin_api"
}
