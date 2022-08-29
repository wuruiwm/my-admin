package task

import (
	"app/config"
	"app/global"
	"app/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

func AdminConfig() {
	adminConfigInit()
	adminConfigWatch()
}

func adminConfigWatch() {
	var (
		sub *redis.PubSub
		err error
	)
	for {
		sub = global.Redis.Subscribe(context.Background(), "admin_config_watch")
		for {
			_, err = sub.ReceiveMessage(context.Background())
			if err != nil {
				break
			}
			adminConfigInit()
		}
		_ = sub.Close()
		time.Sleep(time.Second * 10)
		global.Logger.Error("admin_config_watch", zap.String("error", "连接断开 error:"+err.Error()))
	}
}

func adminConfigInit() {
	configList := make([]*model.AdminConfig, 0)
	global.Db.Table("admin_config").
		Select("group", "key", "value").
		Find(&configList)
	data := make(map[string]map[string]string, 0)
	for _, v := range configList {
		if _, ok := data[v.Group]; !ok {
			data[v.Group] = make(map[string]string, 0)
		}
		data[v.Group][v.Key] = v.Value
	}
	jsonByte, err := json.Marshal(data)
	if err != nil {
		global.Logger.Error("admin_config", zap.Any("configList", configList), zap.Any("data", data), zap.String("error", err.Error()))
		return
	}
	adminConfig := &config.AdminConfig{}
	if err = json.Unmarshal(jsonByte, adminConfig); err != nil {
		global.Logger.Error("admin_config", zap.String("jsonByte", string(jsonByte)), zap.String("error", err.Error()))
		return
	}
	global.Config.AdminConfig = adminConfig
}
