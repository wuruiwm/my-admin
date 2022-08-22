package model

//AdminUser 后台用户表
type AdminUser struct {
	Base
	Username      string `gorm:"column:username;type:char(20);NOT NULL" json:"username"`
	Nickname      string `gorm:"column:nickname;type:char(20);comment:昵称;NOT NULL" json:"nickname"`
	Password      string `gorm:"column:password;type:char(32);comment:hash后的密码;NOT NULL" json:"password"`
	Salt          string `gorm:"column:salt;type:char(36);comment:hash密码的时候加入的salt;NOT NULL" json:"salt"`
	LastLoginTime string `gorm:"column:last_login_time;type:datetime;comment:最后登陆时间" json:"last_login_time"`
	AdminRoleId   string `gorm:"column:admin_role_id;type:char(36);comment:角色id;NOT NULL" json:"admin_role_id"`
}

func (m *AdminUser) TableName() string {
	return "admin_user"
}
