package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"errors"
	"fmt"
)

func AdminApiList(param *request.AdminApiList) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
		where  = make(map[string]interface{}, 0)
	)
	if param.Method != "" {
		where["aa.method"] = param.Method
	}
	if param.Group != "" {
		where["aa.group"] = param.Group
	}
	apiList := make([]*model.AdminApi, 0)
	global.Db.Table("admin_api aa").
		Where(param.GetKeywordWhere("aa.name", "aa.path")).
		Where(param.GetTimeWhere("aa.create_time")).
		Where(where).
		Count(&count)
	if count != 0 {
		global.Db.Table("admin_api aa").
			Where(param.GetKeywordWhere("aa.name", "aa.path")).
			Where(param.GetTimeWhere("aa.create_time")).
			Where(where).
			Order("aa.group asc,aa.path asc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&apiList)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     apiList,
	}, nil
}

func AdminApiCheckUnique(path, method, id string) error {
	where := map[string]interface{}{
		"path":   path,
		"method": method,
	}
	if id != "" {
		where["id <> ?"] = id
	}
	if global.Db.Where(where).
		Select("id").
		Take(&model.AdminApi{}).Error == nil {
		return errors.New(fmt.Sprintf("该api请求已存在 path:%s method:%s", path, method))
	}
	return nil
}

func AdminApiCreate(param *request.AdminApiCreate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminApiCheckUnique(param.Path, param.Method, ""); err != nil {
		return err
	}

	api := &model.AdminApi{
		Name:   param.Name,
		Group:  param.Group,
		Path:   param.Path,
		Method: param.Method,
	}
	if err := global.Db.Create(api).Error; err != nil {
		return errors.New("api创建失败 error: " + err.Error())
	}
	return nil
}

func AdminApiUpdate(param *request.AdminApiUpdate) error {
	if err := param.Check(); err != nil {
		return err
	}
	if err := AdminApiCheckUnique(param.Path, param.Method, param.Id); err != nil {
		return err
	}

	api := &model.AdminApi{
		Name:   param.Name,
		Group:  param.Group,
		Path:   param.Path,
		Method: param.Method,
	}
	if err := global.Db.Where("id", param.Id).
		Select("name", "group", "path", "method").
		Updates(api).Error; err != nil {
		return errors.New("api修改失败 error: " + err.Error())
	}
	return nil
}

func AdminApiDelete(param *request.AdminApiDelete) error {
	if err := global.Db.Where("id", param.Id).
		Delete(&model.AdminApi{}).Error; err != nil {
		return errors.New("api删除失败 error: " + err.Error())
	}
	return nil
}

func AdminApiGroup() []string {
	group := make([]string, 0)
	global.Db.Model(&model.AdminApi{}).
		Order("`group` asc").
		Pluck("distinct `group` as `group`", &group)
	return group
}
