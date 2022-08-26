package model

//Password 密码本表
type Password struct {
	Base
	Name     string `gorm:"column:name;type:char(20);NOT NULL" json:"name"`
	Username string `gorm:"column:username;type:char(50);NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:char(50);NOT NULL" json:"password"`
	Remark   string `gorm:"column:remark;type:varchar(1000);NOT NULL" json:"remark"`
}

func (m *Password) TableName() string {
	return "password"
}
