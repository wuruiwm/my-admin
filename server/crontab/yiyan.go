package crontab

import (
	"app/global"
	"app/model"
	"app/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddieivan01/nic"
	"strings"
	"time"
)

func Yiyan() {
	var success int
	for i := 0; i < 10000; i++ {
		if err := yiyan(); err == nil {
			success++
		}
		time.Sleep(time.Second)
	}
	_ = util.Notice("一言爬虫", fmt.Sprintf("有效新增%d条", success))
}

type yiyanResponse struct {
	Hitokoto string `json:"hitokoto"`
}

func yiyan() error {
	resp, err := nic.Get("https://v1.hitokoto.cn", nil)
	if err != nil {
		return errors.New("一言接口请求失败 error:" + err.Error())
	}
	yiyanResp := &yiyanResponse{}
	if err = json.Unmarshal(resp.Bytes, yiyanResp); err != nil {
		return errors.New("一言json反序列化失败 error:" + err.Error())
	}
	if yiyanResp.Hitokoto == "" {
		return errors.New("一言内容为空")
	}
	if err = global.Db.Create(&model.Yiyan{
		Content: YiyanTransformSymbol(yiyanResp.Hitokoto),
	}).Error; err != nil {
		return errors.New("一言插入数据库 error:" + err.Error())
	}
	return nil
}

func YiyanTransformSymbol(content string) string {
	replace := map[string]string{
		",": "，",
		".": "。",
		"?": "？",
		"!": "！",
	}
	for k, v := range replace {
		content = strings.ReplaceAll(content, k, v)
	}
	return content
}
