package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"context"
	"errors"
)

func AdminRoleList(param *request.AdminRoleList) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
	)
	roleList := make([]*model.AdminRole, 0)
	global.Db.Table("admin_role ar").
		Where(param.GetKeywordWhere("ar.name")).
		Where(param.GetTimeWhere("ar.create_time")).
		Count(&count)
	if count != 0 {
		global.Db.Table("admin_role ar").
			Where(param.GetKeywordWhere("ar.name")).
			Where(param.GetTimeWhere("ar.create_time")).
			Order("ar.update_time desc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&roleList)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     roleList,
	}, nil
}

func AdminRoleCheckUnique(name, id string) error {
	where := map[string]interface{}{
		"name": name,
	}
	if id != "" {
		where["id <> ?"] = id
	}
	if global.Db.Where(where).
		Select("id").
		Take(&model.AdminRole{}).Error == nil {
		return errors.New("角色名称已存在")
	}
	return nil
}

func AdminRoleCreate(param *request.AdminRoleCreate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminRoleCheckUnique(param.Name, ""); err != nil {
		return err
	}

	role := &model.AdminRole{
		Name: param.Name,
	}
	if err := global.Db.Create(role).Error; err != nil {
		return errors.New("角色创建失败 error: " + err.Error())
	}
	return nil
}

func AdminRoleUpdate(param *request.AdminRoleUpdate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminRoleCheckUnique(param.Name, param.Id); err != nil {
		return err
	}

	role := &model.AdminRole{
		Name: param.Name,
	}
	if err := global.Db.Where("id", param.Id).
		Select("name").
		Updates(role).Error; err != nil {
		return errors.New("角色修改失败 error: " + err.Error())
	}
	return nil
}

func AdminRoleDelete(param *request.AdminRoleDelete) error {
	user := &model.AdminUser{}
	if global.Db.Table("admin_user_role aur").
		Joins("inner join admin_user as au on aur.admin_user_id=au.id").
		Where("aur.admin_role_id", param.Id).
		Select("au.id", "au.username").
		Take(user).Error == nil {
		return errors.New("删除失败，该角色有用户正在使用：" + user.Username)
	}
	if err := global.Db.Where("id", param.Id).
		Delete(&model.AdminRole{}).Error; err != nil {
		return errors.New("角色删除失败 error: " + err.Error())
	}
	return nil
}

func AdminRoleAuth(param *request.AdminRoleAuth) *response.AdminRoleAuth {
	menuList := make([]*model.AdminMenu, 0)
	global.Db.Table("admin_menu").
		Select("id", "title as `name`", "parent_id").
		Order("sort asc").
		Find(&menuList)
	apiList := make([]*model.AdminApi, 0)
	global.Db.Table("admin_api").
		Select("id", "name", "group").
		Find(&apiList)
	adminRoleAuth := &response.AdminRoleAuth{
		Menu: make([]*response.AdminRoleAuthDetail, 0),
		Api:  make([]*response.AdminRoleAuthDetail, 0),
	}
	for _, v := range menuList {
		adminRoleAuthMenu := &response.AdminRoleAuthDetail{
			ParentId: v.ParentId,
		}
		adminRoleAuthMenu.Id = v.Id
		adminRoleAuthMenu.Name = v.Name
		adminRoleAuth.Menu = append(adminRoleAuth.Menu, adminRoleAuthMenu)
	}
	for _, v := range apiList {
		adminRoleAuth.Api = append(adminRoleAuth.Api, &response.AdminRoleAuthDetail{
			Id:       v.Id,
			Name:     v.Name,
			ParentId: v.Group,
		})
	}
	roleMenuListTmp := make([]*model.AdminRoleMenu, 0)
	global.Db.Table("admin_role_menu").
		Where("admin_role_id", param.Id).
		Select("admin_menu_id").
		Find(&roleMenuListTmp)
	roleMenuList := make(map[string]int, 0)
	for _, v := range roleMenuListTmp {
		roleMenuList[v.AdminMenuId] = 1
	}
	roleApiListTmp := make([]*model.AdminRoleApi, 0)
	global.Db.Table("admin_role_api").
		Where("admin_role_id", param.Id).
		Select("admin_api_id").
		Find(&roleApiListTmp)
	roleApiList := make(map[string]int, 0)
	for _, v := range roleApiListTmp {
		roleApiList[v.AdminApiId] = 1
	}
	for k, v := range adminRoleAuth.Menu {
		adminRoleAuth.Menu[k].IsChecked = roleMenuList[v.Id]
	}
	for k, v := range adminRoleAuth.Api {
		adminRoleAuth.Api[k].IsChecked = roleApiList[v.Id]
	}
	return adminRoleAuth
}

func AdminRoleAuthUpdate(param *request.AdminRoleAuthUpdate) error {
	apiList := make([]*model.AdminRoleApi, 0)
	menuList := make([]*model.AdminRoleMenu, 0)
	for _, v := range param.ApiId {
		apiList = append(apiList, &model.AdminRoleApi{
			AdminRoleId: param.Id,
			AdminApiId:  v,
		})
	}
	for _, v := range param.MenuId {
		menuList = append(menuList, &model.AdminRoleMenu{
			AdminRoleId: param.Id,
			AdminMenuId: v,
		})
	}
	tx := global.Db.Begin()
	if err := tx.Where("admin_role_id", param.Id).
		Delete(&model.AdminRoleApi{}).Error; err != nil {
		tx.Rollback()
		return errors.New("更新权限失败 error: " + err.Error())
	}
	if err := tx.Where("admin_role_id", param.Id).
		Delete(&model.AdminRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return errors.New("更新权限失败 error: " + err.Error())
	}
	if len(apiList) != 0 {
		if err := tx.Create(&apiList).Error; err != nil {
			tx.Rollback()
			return errors.New("更新权限失败 error: " + err.Error())
		}
	}
	if len(menuList) != 0 {
		if err := tx.Create(&menuList).Error; err != nil {
			tx.Rollback()
			return errors.New("更新权限失败 error: " + err.Error())
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("更新权限失败 error: " + err.Error())
	}

	if err := global.Redis.Del(context.Background(), "admin_role_api_auth:"+param.Id).
		Err(); err != nil {
		return errors.New("删除权限缓存失败 error:" + err.Error())
	}
	return nil
}
