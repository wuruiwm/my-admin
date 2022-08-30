package logic

import (
	"app/api/request"
	"app/global"
	"app/model"
	"errors"
	"fmt"
	"math/rand"
)

func Yiyan(param *request.Yiyan) (interface{}, error) {
	var count int64
	global.Db.Model(&model.Yiyan{}).Count(&count)
	if count == 0 {
		return nil, errors.New("一言条数为0")
	}
	offset := rand.Intn(int(count))
	yiyan := &model.Yiyan{}
	if err := global.Db.Order("create_time asc").
		Offset(offset).
		Limit(1).
		Take(yiyan).Error; err != nil {
		return nil, errors.New("一言不存在 error:" + err.Error())
	}
	if param.ResponseType == "js" {
		return fmt.Sprintf(`function hitokoto(){document.write("%s");}`, yiyan.Content), nil
	} else if param.ResponseType == "text" {
		return yiyan.Content, nil
	} else {
		return yiyan, nil
	}
}
