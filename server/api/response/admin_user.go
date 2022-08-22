package response

//AdminUserLogin 后台用户登录返回
type AdminUserLogin struct {
	Id          string `json:"id"`
	AdminRoleId string `json:"admin_role_id"`
	ExpireTime  string `json:"expire_time"`
	Token       string `json:"token"`
}

type AdminUserInfo struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	AdminRoleId   string `json:"admin_role_id"`
	AdminRoleName string `json:"admin_role_name"`
}
