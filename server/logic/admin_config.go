package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"app/util"
	"errors"
)

func AdminConfigList(param *request.AdminConfigList) []*response.AdminConfig {
	configList := make([]*model.AdminConfig, 0)
	global.Db.Where("group", param.Group).
		Select("key", "value").
		Find(&configList)
	adminConfigList := make([]*response.AdminConfig, 0)
	for _, v := range configList {
		adminConfigList = append(adminConfigList, &response.AdminConfig{
			Key:   v.Key,
			Value: v.Value,
		})
	}
	return adminConfigList
}

func AdminConfigUpdate(param *request.AdminConfigUpdate) error {
	configList := make([]*model.AdminConfig, 0)
	for _, v := range param.Data {
		configList = append(configList, &model.AdminConfig{
			Group: param.Group,
			Key:   v.Key,
			Value: v.Value,
		})
	}
	tx := global.Db.Begin()
	if err := tx.Where("group", param.Group).
		Delete(&model.AdminConfig{}).Error; err != nil {
		tx.Rollback()
		return errors.New("更新配置失败 error: " + err.Error())
	}
	if len(configList) != 0 {
		if err := tx.Create(&configList).Error; err != nil {
			tx.Rollback()
			return errors.New("更新配置失败 error: " + err.Error())
		}
	}
	if err := tx.Commit().Error; err != nil {
		return errors.New("更新配置失败 error: " + err.Error())
	}

	mq, err := util.NewRabbitmq()
	if err != nil {
		return errors.New("删除配置缓存失败 error:" + err.Error())
	}
	if err = mq.Publish("admin_config", "fanout", "1", 0); err != nil {
		return errors.New("删除配置缓存失败 error:" + err.Error())
	}
	return nil
}
