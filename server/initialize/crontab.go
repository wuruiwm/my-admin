package initialize

import (
	"app/global"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

func CronTab() {
	//添加定时任务 cron表达式文档 https://pkg.go.dev/github.com/robfig/cron
	c := cron.New(cron.WithSeconds())
	cronTabList := []*cronTab{
		//NewCronTab("second", "* * * * * *", crontab.Second, 1),  //秒  级 定时任务
		//NewCronTab("minute", "0 * * * * *", crontab.Minute, 60), //分钟级 定时任务
	}
	for _, v := range cronTabList {
		_, err := c.AddFunc(v.spec, v.run)
		if err != nil {
			panic(err)
		}
		c.Start()
	}
}

type cronTab struct {
	name       string //定时任务分布式锁 保证唯一
	spec       string
	task       func()
	maxRunTime int //最大执行时间(秒) 如果还没有释放锁 则其他定时任务可重新开始执行
}

func NewCronTab(name string, spec string, task func(), maxRunTime int) *cronTab {
	return &cronTab{
		name:       name,
		spec:       spec,
		task:       task,
		maxRunTime: maxRunTime,
	}
}

func (c *cronTab) run() {
	ok, err := global.Redis.SetNX(context.Background(), "crontab_lock:"+c.name, 1, time.Duration(c.maxRunTime)*time.Second).Result()
	if err != nil {
		global.Logger.Error("crontab_lock", zap.Any("cronTab", c), zap.String("error", "定时任务分布式锁 error:"+err.Error()))
	}
	if ok {
		c.task()
	}
}
