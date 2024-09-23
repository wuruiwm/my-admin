package initialize

import (
	"app/crontab"
	"app/global"
	"app/util"
	"context"
	"github.com/robfig/cron/v3"
	"time"
)

func Crontab() {
	//添加定时任务 cron表达式文档 https://pkg.go.dev/github.com/robfig/cron
	c := cron.New()
	cronTabList := []*CronTabTask{
		NewCrontab("yiyan", "0 0 * * 1", crontab.Yiyan, 3600*8),
		NewCrontab("cloudflare", "1,5,10,15,20,30 8,18 * * *", crontab.Cloudflare, 60*10),
		NewCrontab("openwrt", "* * * * *", crontab.Openwrt, 6),
		NewCrontab("aliyunCdt", "*/10 * * * *", crontab.AliyunCdt, 800),
	}
	for _, v := range cronTabList {
		_, err := c.AddFunc(v.spec, v.run)
		if err != nil {
			panic(err)
		}
		c.Start()
	}
}

type CronTabTask struct {
	name       string //定时任务分布式锁 保证唯一
	spec       string
	task       func()
	maxRunTime int //最大执行时间(秒) 如果还没有释放锁 则其他定时任务可重新开始执行
}

func NewCrontab(name string, spec string, task func(), maxRunTime int) *CronTabTask {
	return &CronTabTask{
		name:       name,
		spec:       spec,
		task:       task,
		maxRunTime: maxRunTime,
	}
}

func (c *CronTabTask) run() {
	key := "crontab_lock:" + c.name
	ok, err := global.Redis.SetNX(context.Background(), key, 1, time.Duration(c.maxRunTime)*time.Second).Result()
	if err != nil {
		util.Logger.Error("crontab_lock", util.Map{
			"name":       c.name,
			"spec":       c.spec,
			"maxRunTime": c.maxRunTime,
			"error":      "定时任务分布式锁lock error:" + err.Error(),
		})
	}
	if ok {
		c.task()
		if err = global.Redis.Del(context.Background(), key).Err(); err != nil {
			util.Logger.Error("crontab_lock", util.Map{
				"name":       c.name,
				"spec":       c.spec,
				"maxRunTime": c.maxRunTime,
				"error":      "定时任务分布式锁unlock error:" + err.Error(),
			})
		}
	}
}
