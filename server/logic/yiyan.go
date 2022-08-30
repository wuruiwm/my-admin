package logic

import (
	"app/api/request"
	"app/global"
	"app/model"
	"errors"
)

func Yiyan(param *request.Yiyan) (interface{}, error) {
	var count int64
	global.Db.Model(&model.Yiyan{}).Count(&count)
	if count == 0 {
		return nil, errors.New("一言条数为0")
	}

	return nil, nil
}
