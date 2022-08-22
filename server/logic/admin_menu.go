package logic

import (
	"app/api/request"
	"app/global"
	"app/model"
	"app/util"
	"errors"
	"fmt"
)

func AdminMenuList() []*model.AdminMenu {
	menuList := make([]*model.AdminMenu, 0)
	global.Db.Table("admin_menu").
		Order("is_hidden asc,sort asc").
		Find(&menuList)
	return menuList
}

func AdminMenuCheckUnique(name, id string) error {
	where := map[string]interface{}{
		"name": name,
	}
	if id != "" {
		where["id <> ?"] = id
	}
	if global.Db.Where(where).
		Select("id").
		Take(&model.AdminMenu{}).Error == nil {
		return errors.New(fmt.Sprintf("该路由name已存在 name:%s", name))
	}
	return nil
}

func AdminMenuCheckParentMenu(parentId string) error {
	if parentId != "" && global.Db.Where("id", parentId).
		Select("id").
		Take(&model.AdminMenu{}).Error != nil {
		return errors.New("上级菜单不存在")
	}
	return nil
}

func AdminMenuCreate(param *request.AdminMenuCreate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminMenuCheckUnique(param.Name, ""); err != nil {
		return err
	}
	if err := AdminMenuCheckParentMenu(param.ParentId); err != nil {
		return err
	}

	maxSort := &model.AdminMenu{}
	global.Db.Model(&model.AdminMenu{}).Select("max(sort) as `sort`").Scan(maxSort)
	maxSort.Sort++
	menu := &model.AdminMenu{
		Path:      param.Path,
		Name:      param.Name,
		ParentId:  param.ParentId,
		Component: param.Component,
		Icon:      param.Icon,
		Title:     param.Title,
		Sort:      maxSort.Sort,
		IsHidden:  param.IsHidden,
	}
	if err := global.Db.Create(menu).Error; err != nil {
		return errors.New("菜单创建失败 error: " + err.Error())
	}
	return nil
}

func AdminMenuUpdate(param *request.AdminMenuUpdate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminMenuCheckUnique(param.Name, param.Id); err != nil {
		return err
	}
	if err := AdminMenuCheckParentMenu(param.ParentId); err != nil {
		return err
	}

	menu := &model.AdminMenu{
		Path:      param.Path,
		Name:      param.Name,
		ParentId:  param.ParentId,
		Component: param.Component,
		Icon:      param.Icon,
		Title:     param.Title,
		IsHidden:  param.IsHidden,
	}
	if err := global.Db.Where("id", param.Id).
		Select("name", "parent_id", "component", "icon", "title", "is_hidden").
		Updates(menu).Error; err != nil {
		return errors.New("菜单创建失败 error: " + err.Error())
	}
	return nil
}

func AdminMenuDelete(param *request.AdminMenuDelete) error {
	if err := global.Db.Where("id", param.Id).
		Delete(&model.AdminMenu{}).Error; err != nil {
		return errors.New("菜单删除失败 error: " + err.Error())
	}
	return nil
}

func AdminMenuSort(param *request.AdminMenuSort) error {
	if !util.InArray(param.Type, []int{0, 1}) {
		return errors.New("type不正确")
	}
	oldMenu := &model.AdminMenu{}
	if global.Db.Where("id", param.Id).
		Select("parent_id", "sort").
		Take(oldMenu).Error != nil {
		return errors.New("该菜单不存在")
	}
	where := map[string]interface{}{
		"is_hidden": 0,
		"parent_id": oldMenu.ParentId,
	}
	var (
		sortWhere string
		sortOrder string
	)
	if param.Type == 0 {
		sortWhere = "sort < ?"
		sortOrder = "sort desc"
	} else if param.Type == 1 {
		sortWhere = "sort > ?"
		sortOrder = "sort asc"
	}
	newMenu := &model.AdminMenu{}
	if global.Db.Where(where).
		Where(sortWhere, oldMenu.Sort).
		Select("id", "sort").
		Order(sortOrder).
		Take(newMenu).Error != nil {
		return errors.New("操作异常")
	}
	tx := global.Db.Begin()
	if err := tx.Where("id", newMenu.Id).
		Select("sort").
		Updates(&model.AdminMenu{Sort: 0}).Error; err != nil {
		tx.Rollback()
		return errors.New("操作失败 error:" + err.Error())
	}
	if err := tx.Where("id", param.Id).
		Select("sort").
		Updates(&model.AdminMenu{Sort: newMenu.Sort}).Error; err != nil {
		tx.Rollback()
		return errors.New("操作失败 error:" + err.Error())
	}
	if err := tx.Where("id", newMenu.Id).
		Select("sort").
		Updates(&model.AdminMenu{Sort: oldMenu.Sort}).Error; err != nil {
		tx.Rollback()
		return errors.New("操作失败 error:" + err.Error())
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("操作失败 error: " + err.Error())
	}
	return nil
}
