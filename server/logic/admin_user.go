package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"app/util"
	"errors"
	"gorm.io/gorm"
	"strings"
	"unicode/utf8"
)

func AdminUserAuth(token, path, method string) (*response.AdminUserLogin, int, error) {
	adminUserLogin, err := AdminUserLoginAuth(token)
	if err != nil {
		return nil, response.LoginExpiredCode, err
	}
	if err = AdminUserApiAuth(adminUserLogin.AdminRoleId, path, method); err != nil {
		return nil, response.ErrorCode, err
	}
	return adminUserLogin, response.SuccessCode, nil
}

func AdminUserLoginAuth(token string) (*response.AdminUserLogin, error) {
	if token == "" {
		return nil, errors.New("未登录")
	}
	adminUserLogin := &response.AdminUserLogin{}
	cacheKey := "admin_token:" + token
	if err := util.GetCache(cacheKey, adminUserLogin); err != nil {
		return nil, errors.New("登录已过期")
	}
	expireTime, err := util.DateToUnix(adminUserLogin.ExpireTime)
	if err != nil {
		return nil, errors.New("登录有效期异常,请重新登录")
	}
	if expireTime < util.GetUnix() {
		return nil, errors.New("登录已过期")
	}
	adminUser := &model.AdminUser{}
	if global.Db.Where("id = ?", adminUserLogin.Id).
		Select("id", "password", "salt", "last_login_time", "admin_role_id").
		Take(adminUser).Error != nil {
		return nil, errors.New("该用户不存在")
	}
	if adminUserLogin.Token != AdminUserHashToken(adminUser, adminUser.LastLoginTime) {
		return nil, errors.New("登录失效")
	}
	return adminUserLogin, nil
}

func AdminUserApiAuth(userRoleId string, urlPath string, method string) error {
	var err error
	urlPath = strings.ToLower(urlPath)
	method = strings.ToLower(method)
	result := make(map[string]int, 0)
	//获取缓存的api权限
	cacheKey := "admin_role_api_auth:" + userRoleId
	if err = util.GetCache(cacheKey, &result); err != nil {
		apiList := make([]*model.AdminApi, 0)
		global.Db.Table("admin_role as ar").
			Joins("inner join admin_role_api as ara on ar.id=ara.admin_role_id").
			Joins("inner join admin_api as aa on ara.admin_api_id=aa.id").
			Where("ar.id", userRoleId).
			Select("aa.method", "aa.path").
			Find(&apiList)
		for _, v := range apiList {
			result[strings.ToLower(v.Path)+"_"+strings.ToLower(v.Method)] = 1
		}
		err = util.SetCache(cacheKey, result, 0)
		if err != nil {
			return errors.New("角色api权限缓存失败 error:" + err.Error())
		}
	}
	if _, ok := result[urlPath+"_"+method]; !ok {
		return errors.New("无权限")
	}
	return nil
}

func AdminUserCreateToken(adminUser *model.AdminUser) (*response.AdminUserLogin, error) {
	time := util.GetUnix()
	lastLoginTime := util.UnixToDate(time)
	expireTime := util.UnixToDate(time + 60*60*24*30)
	if err := AdminUserUpdateLastLoginTime(global.Db, adminUser.Id, lastLoginTime); err != nil {
		return nil, errors.New("更新用户最后登录时间失败 error: " + err.Error())
	}

	adminUserLogin := &response.AdminUserLogin{
		Id:          adminUser.Id,
		AdminRoleId: adminUser.AdminRoleId,
		ExpireTime:  expireTime,
		Token:       AdminUserHashToken(adminUser, lastLoginTime),
	}
	cacheKey := "admin_token:" + adminUserLogin.Token
	if err := util.SetCache(cacheKey, adminUserLogin, 60*60*24*30); err != nil {
		return nil, errors.New("保存token失败 error: " + err.Error())
	}
	return adminUserLogin, nil
}

func AdminUserHashToken(adminUser *model.AdminUser, lastLoginTime string) string {
	return util.MD5(adminUser.Id + adminUser.AdminRoleId + adminUser.Password + adminUser.Salt + lastLoginTime)
}

func AdminUserLogin(param *request.AdminUserLogin) (*response.AdminUserLogin, error) {
	adminUser, err := AdminUserCheckPassword(param.Username, param.Password, "id", "password", "salt", "admin_role_id")
	if err != nil {
		return nil, err
	}
	return AdminUserCreateToken(adminUser)
}

func AdminUserHashPassword(password string, salt string) string {
	return util.MD5(util.MD5(password+global.Config.Key) + salt)
}

func AdminUserUpdateLastLoginTime(db *gorm.DB, id string, lastLoginTime string) error {
	user := &model.AdminUser{
		LastLoginTime: lastLoginTime,
	}
	return db.Where("id = ?", id).
		Select("last_login_time").
		Updates(user).Error
}

func AdminUserMenu(userId string) ([]*model.AdminMenu, error) {
	menuList := make([]*model.AdminMenu, 0)
	err := global.Db.Table("admin_user as au").
		Joins("inner join admin_role as ar on au.admin_role_id=ar.id").
		Joins("inner join admin_role_menu as arm on ar.id=arm.admin_role_id").
		Joins("inner join admin_menu as am on arm.admin_menu_id=am.id").
		Where("au.id", userId).
		Select("am.*").
		Order("am.sort asc").
		Find(&menuList).Error
	return menuList, err
}

func AdminUserList(param *request.AdminUserList) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
		where  = make(map[string]interface{}, 0)
	)
	if param.AdminRoleId != "" {
		where["au.admin_role_id"] = param.AdminRoleId
	}
	userList := make([]*model.AdminUser, 0)
	global.Db.Table("admin_user au").
		Where(where).
		Where(param.GetTimeWhere("au.last_login_time")).
		Where(param.GetKeywordWhere("au.username")).
		Count(&count)
	if count != 0 {
		global.Db.Table("admin_user au").
			Select("au.id", "au.nickname", "au.create_time", "au.update_time", "au.username", "au.last_login_time", "au.admin_role_id").
			Where(where).
			Where(param.GetTimeWhere("au.last_login_time")).
			Where(param.GetKeywordWhere("au.username")).
			Order("au.update_time desc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&userList)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     userList,
	}, nil
}

func AdminUserCheckUnique(username, id string) error {
	where := map[string]interface{}{
		"username": username,
	}
	if id != "" {
		where["id <> ?"] = id
	}
	if global.Db.Where(where).
		Select("id").
		Take(&model.AdminUser{}).Error == nil {
		return errors.New("角色名称已存在")
	}
	return nil
}

func AdminUserCheckRoleExist(adminRoleId string) error {
	if global.Db.Where("id", adminRoleId).
		Select("id").
		Take(&model.AdminRole{}).Error != nil {
		return errors.New("该角色不存在")
	}
	return nil
}

func AdminUserCreate(param *request.AdminUserCreate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if param.Password == "" {
		return errors.New("请填写密码")
	}
	if err := AdminUserCheckUnique(param.Username, ""); err != nil {
		return err
	}
	if err := AdminUserCheckRoleExist(param.AdminRoleId); err != nil {
		return err
	}

	salt := util.Uuid()
	user := &model.AdminUser{
		Username:      param.Username,
		Nickname:      param.Nickname,
		Password:      AdminUserHashPassword(param.Password, salt),
		AdminRoleId:   param.AdminRoleId,
		Salt:          salt,
		LastLoginTime: util.GetDate(),
	}
	if err := global.Db.Create(user).Error; err != nil {
		return errors.New("用户创建失败 error: " + err.Error())
	}
	return nil
}

func AdminUserUpdate(param *request.AdminUserUpdate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminUserCheckUnique(param.Username, param.Id); err != nil {
		return err
	}
	if err := AdminUserCheckRoleExist(param.AdminRoleId); err != nil {
		return err
	}

	updateSelect := []string{"username", "nickname", "admin_role_id"}
	user := &model.AdminUser{
		Username:    param.Username,
		Nickname:    param.Nickname,
		AdminRoleId: param.AdminRoleId,
	}
	tx := global.Db.Begin()
	if err := tx.Where("id", param.Id).
		Select(updateSelect).
		Updates(user).Error; err != nil {
		return errors.New("用户修改失败 error: " + err.Error())
	}
	//密码有值 才修改密码
	if param.Password != "" {
		if err := AdminUserPasswordUpdate(param.Id, param.Password, tx); err != nil {
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("用户修改失败 error: " + err.Error())
	}
	return nil
}

func AdminUserDelete(param *request.AdminUserDelete) error {
	if err := global.Db.Where("id", param.Id).Delete(&model.AdminUser{}).Error; err != nil {
		return errors.New("用户删除失败 error: " + err.Error())
	}
	return nil
}

func AdminUserRole() []*model.AdminRole {
	roleList := make([]*model.AdminRole, 0)
	global.Db.Table("admin_role ar").
		Order("ar.update_time desc").
		Find(&roleList)
	return roleList
}

func AdminUserRoleUpdate(param *request.AdminUserRoleUpdate) error {
	if err := AdminUserCheckRoleExist(param.AdminRoleId); err != nil {
		return err
	}
	user := &model.AdminUser{
		AdminRoleId: param.AdminRoleId,
	}
	if err := global.Db.Where("id", param.Id).
		Select("admin_role_id").
		Updates(user).Error; err != nil {
		return errors.New("用户角色修改失败 error: " + err.Error())
	}
	return nil
}

func AdminUserPasswordUpdate(id, password string, db *gorm.DB) error {
	if password == "" {
		return errors.New("请填写密码")
	}
	if utf8.RuneCountInString(password) < 6 {
		return errors.New("密码不能少于六位")
	}
	salt := util.Uuid()
	user := &model.AdminUser{
		Password: AdminUserHashPassword(password, salt),
		Salt:     salt,
	}
	if err := db.Where("id", id).
		Select("password", "salt").
		Updates(user).Error; err != nil {
		return errors.New("用户密码修改失败 error: " + err.Error())
	}
	return nil
}

func AdminUserInfo(id string) (*response.AdminUserInfo, error) {
	adminUserInfo := &response.AdminUserInfo{}
	if global.Db.Model(&model.AdminUser{}).
		Where("id", id).
		Select("id", "username", "nickname", "admin_role_id").
		Take(adminUserInfo).Error != nil {
		return nil, errors.New("该用户不存在")
	}
	if global.Db.Model(&model.AdminRole{}).
		Where("id", adminUserInfo.AdminRoleId).
		Select("name as admin_role_name").
		Take(adminUserInfo).Error != nil {
		return nil, errors.New("该角色不存在")
	}
	return adminUserInfo, nil
}

func AdminUserCheckPassword(idOrUsername, password string, field ...string) (*model.AdminUser, error) {
	if len(field) == 0 {
		field = []string{"*"}
	}
	user := &model.AdminUser{}
	if global.Db.Where("(id = ? or username = ?)", idOrUsername, idOrUsername).Select(field).Take(user).Error != nil {
		return nil, errors.New("该用户不存在")
	}
	if AdminUserHashPassword(password, user.Salt) != user.Password {
		return user, errors.New("密码不正确")
	}
	return user, nil
}

func AdminUserInfoUpdate(param *request.AdminUserInfoUpdate, id string) error {
	if param.Nickname == "" {
		return errors.New("请填写昵称")
	}
	if param.NewPassword != "" && utf8.RuneCountInString(param.NewPassword) < 6 {
		return errors.New("密码不能少于六位")
	}
	if _, err := AdminUserCheckPassword(id, param.OldPassword, "salt", "password"); err != nil {
		return err
	}

	user := &model.AdminUser{
		Nickname: param.Nickname,
	}
	tx := global.Db.Begin()
	if err := tx.Where("id", id).
		Select("nickname").
		Updates(user).Error; err != nil {
		return errors.New("用户信息修改失败 error: " + err.Error())
	}
	//密码有值 才修改密码
	if param.NewPassword != "" {
		if err := AdminUserPasswordUpdate(id, param.NewPassword, tx); err != nil {
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("用户信息修改失败 error: " + err.Error())
	}
	return nil
}
