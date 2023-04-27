package model

// Youtube youtube下载表
type Youtube struct {
	Base
	Name    string `gorm:"column:name;type:char(20);comment:音乐名;NOT NULL" json:"name"`
	Url     string `gorm:"column:url;type:varchar(100);comment:youtube视频地址;NOT NULL" json:"url"`
	Status  uint   `gorm:"column:status;type:tinyint(1) unsigned;default:0;comment:状态 0待执行 1执行成功 2执行失败;NOT NULL" json:"status"`
	Command string `gorm:"column:command;type:char(200);comment:执行的命令;NOT NULL" json:"command"`
	Content string `gorm:"column:content;type:text;comment:命令执行日志;NOT NULL" json:"content"`
}

func (m *Youtube) TableName() string {
	return "youtube"
}
