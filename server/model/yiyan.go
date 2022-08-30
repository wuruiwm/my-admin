package model

//Yiyan 一言表
type Yiyan struct {
	Base
	Content string `gorm:"column:content;type:varchar(255);NOT NULL" json:"content"`
}

func (m *Yiyan) TableName() string {
	return "yiyan"
}
