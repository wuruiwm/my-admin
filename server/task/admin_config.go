package task

import (
	"app/config"
	"app/global"
	"app/model"
	"app/util"
	"github.com/bytedance/sonic"
	"github.com/rabbitmq/amqp091-go"
)

func AdminConfig() {
	adminConfigInit()
	go adminConfigWatch()
}

func adminConfigWatch() {
	mq, err := util.NewRabbitmq()
	if err != nil {
		panic(err)
	}
	mq.Consume("admin_config", "fanout", func(delivery amqp091.Delivery, rabbitmq *util.Rabbitmq) {
		adminConfigInit()
	})
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
	jsonByte, err := sonic.Marshal(data)
	if err != nil {
		util.NewLogger().Error("admin_config", util.Map{
			"configList": configList,
			"data":       data,
			"error":      err.Error(),
		})
		return
	}
	adminConfig := &config.AdminConfig{}
	if err = sonic.Unmarshal(jsonByte, adminConfig); err != nil {
		util.NewLogger().Error("admin_config", util.Map{
			"jsonByte": string(jsonByte),
			"error":    err.Error(),
		})
		return
	}
	global.Config.AdminConfig = adminConfig
}
