package model

//AdminMenu 后台菜单表
type AdminMenu struct {
	Base
	Path      string `gorm:"column:path;type:char(20);comment:路由path;NOT NULL" json:"path"`
	Name      string `gorm:"column:name;type:char(20);comment:路由name;NOT NULL" json:"name"`
	ParentId  string `gorm:"column:parent_id;type:char(36);comment:父菜单id;NOT NULL" json:"parent_id"`
	Component string `gorm:"column:component;type:char(255);comment:文件路径;NOT NULL" json:"component"`
	Icon      string `gorm:"column:icon;type:char(20);comment:菜单图标;NOT NULL" json:"icon"`
	Title     string `gorm:"column:title;type:char(10);comment:菜单名称;NOT NULL" json:"title"`
	Sort      uint   `gorm:"column:sort;type:tinyint(4) unsigned;default:0;comment:排序 越大排序越靠前;NOT NULL" json:"sort"`
	IsHidden  uint   `gorm:"column:is_hidden;type:tinyint(1) unsigned;default:0;comment:是否隐藏菜单 0显示 1隐藏 ;NOT NULL" json:"is_hidden"`
}

func (m *AdminMenu) TableName() string {
	return "admin_menu"
}
