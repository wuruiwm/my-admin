package logic

import (
	"app/api/request"
	"app/api/response"
	"app/global"
	"app/model"
	"app/util"
	"errors"
	"github.com/bytedance/sonic"
)

func YoutubeCreate(param *request.YoutubeCreate) error {
	if err := param.Check(); err != nil {
		return err
	}
	youtube := &model.Youtube{
		Name: param.Name,
		Url:  param.Url,
	}
	if err := global.Db.Create(youtube).Error; err != nil {
		return errors.New("youtube下载创建失败 error: " + err.Error())
	}
	return youtubePublish(youtube)
}

func YoutubeList(param *request.YoutubeList) (*response.Page, error) {
	var (
		count  int64
		offset = (param.Page - 1) * param.PageSize
	)
	youtubeList := make([]*model.Youtube, 0)
	global.Db.Table("youtube").
		Where(param.GetKeywordWhere("name")).
		Where(param.GetTimeWhere("create_time")).
		Count(&count)
	if count != 0 {
		global.Db.Table("youtube").
			Where(param.GetKeywordWhere("name")).
			Where(param.GetTimeWhere("create_time")).
			Order("create_time desc").
			Offset(offset).
			Limit(param.PageSize).
			Find(&youtubeList)
	}
	return &response.Page{
		Page:     param.Page,
		PageSize: param.PageSize,
		Total:    count,
		List:     youtubeList,
	}, nil
}

func YoutubeRetry(param *request.YoutubeRetry) error {
	youtube := &model.Youtube{
		Status:  0,
		Command: "",
		Content: "",
	}
	if err := global.Db.Where("id", param.Id).
		Select("status", "command", "content").
		Updates(youtube).Error; err != nil {
		return errors.New("修改失败 error: " + err.Error())
	}
	if err := global.Db.Where("id", param.Id).Find(youtube).Error; err != nil {
		return errors.New("数据异常 error: " + err.Error())
	}
	return youtubePublish(youtube)
}

func YoutubeDelete(param *request.YoutubeDelete) error {
	if err := global.Db.Where("id", param.Id).
		Delete(&model.Youtube{}).Error; err != nil {
		return errors.New("删除失败 error: " + err.Error())
	}
	return nil
}

func youtubePublish(youtube *model.Youtube) error {
	json, err := sonic.Marshal(youtube)
	if err != nil {
		return errors.New("json error: " + err.Error())
	}
	mq, err := util.NewRabbitmq()
	if err != nil {
		return errors.New("队列连接失败 error: " + err.Error())
	}
	err = mq.Publish("youtube", "direct", json, 0)
	if err != nil {
		return errors.New("队列插入失败 error: " + err.Error())
	}
	return nil
}
