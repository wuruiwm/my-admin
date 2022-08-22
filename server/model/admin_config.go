package model

//AdminConfig 后台配置表
type AdminConfig struct {
	Base
	Group string `gorm:"column:group;type:char(20);comment:分组;NOT NULL" json:"group"`
	Key   string `gorm:"column:key;type:char(50);comment:配置key;NOT NULL" json:"key"`
	Value string `gorm:"column:value;type:text;comment:配置value;NOT NULL" json:"value"`
}

func (m *AdminConfig) TableName() string {
	return "admin_config"
}
