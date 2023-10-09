package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"app/util"
	"errors"
	"github.com/bytedance/sonic"
	"strings"
)

func M3u8Create(param *request.M3u8Create) error {
	if err := param.Check(); err != nil {
		return err
	}
	array := strings.Split(param.Url, "&")

	m3u8 := &model.M3u8{
		Name: param.Name,
		Url:  array[0],
	}
	if err := global.Db.Create(m3u8).Error; err != nil {
		return errors.New("m3u8下载创建失败 error: " + err.Error())
	}
	return m3u8Publish(m3u8)
}

func M3u8List(param *request.M3u8List) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
	)
	m3u8List := make([]*model.M3u8, 0)
	global.Db.Table("m3u8").
		Where(param.GetKeywordWhere("name")).
		Where(param.GetTimeWhere("create_time")).
		Count(&count)
	if count != 0 {
		global.Db.Table("m3u8").
			Where(param.GetKeywordWhere("name")).
			Where(param.GetTimeWhere("create_time")).
			Order("create_time desc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&m3u8List)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     m3u8List,
	}, nil
}

func M3u8Retry(param *request.M3u8Retry) error {
	m3u8 := &model.M3u8{
		Status:  0,
		Command: "",
		Content: "",
	}
	if err := global.Db.Where("id", param.Id).
		Select("status", "command", "content").
		Updates(m3u8).Error; err != nil {
		return errors.New("修改失败 error: " + err.Error())
	}
	if err := global.Db.Where("id", param.Id).Find(m3u8).Error; err != nil {
		return errors.New("数据异常 error: " + err.Error())
	}
	return m3u8Publish(m3u8)
}

func M3u8Delete(param *request.M3u8Delete) error {
	if err := global.Db.Where("id", param.Id).
		Delete(&model.M3u8{}).Error; err != nil {
		return errors.New("删除失败 error: " + err.Error())
	}
	return nil
}

func m3u8Publish(m3u8 *model.M3u8) error {
	json, err := sonic.Marshal(m3u8)
	if err != nil {
		return errors.New("json error: " + err.Error())
	}
	mq, err := util.NewRabbitmq()
	if err != nil {
		return errors.New("队列连接失败 error: " + err.Error())
	}
	err = mq.Publish("m3u8", "direct", json, 0)
	if err != nil {
		return errors.New("队列插入失败 error: " + err.Error())
	}
	return nil
}
