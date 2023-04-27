package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"errors"
)

func PasswordList(param *request.PasswordList) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
	)
	passwordList := make([]*model.Password, 0)
	global.Db.Table("password").
		Where(param.GetKeywordWhere("name", "username", "remark")).
		Where(param.GetTimeWhere("create_time")).
		Count(&count)
	if count != 0 {
		global.Db.Table("password").
			Where(param.GetKeywordWhere("name", "username", "remark")).
			Where(param.GetTimeWhere("create_time")).
			Order("create_time desc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&passwordList)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     passwordList,
	}, nil
}

func PasswordCreate(param *request.PasswordCreate) error {
	if err := param.Check(); err != nil {
		return err
	}

	password := &model.Password{
		Name:     param.Name,
		Username: param.Username,
		Password: param.Password,
		Remark:   param.Remark,
	}
	if err := global.Db.Create(password).Error; err != nil {
		return errors.New("密码创建失败 error: " + err.Error())
	}
	return nil
}

func PasswordUpdate(param *request.PasswordUpdate) error {
	if err := param.Check(); err != nil {
		return err
	}

	password := &model.Password{
		Name:     param.Name,
		Username: param.Username,
		Password: param.Password,
		Remark:   param.Remark,
	}
	if err := global.Db.Where("id", param.Id).
		Select("name", "username", "password", "remark").
		Updates(password).Error; err != nil {
		return errors.New("密码修改失败 error: " + err.Error())
	}
	return nil
}

func PasswordDelete(param *request.PasswordDelete) error {
	if err := global.Db.Where("id", param.Id).
		Delete(&model.Password{}).Error; err != nil {
		return errors.New("密码删除失败 error: " + err.Error())
	}
	return nil
}
