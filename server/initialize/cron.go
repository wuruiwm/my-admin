package initialize

import (
	"app/cron"
	"app/util"
	crontab "github.com/robfig/cron/v3"
)

type CronList struct {
	spec string
	run  func()
}

func Cron() {
	util.NoticeGotify("测试标题", "测试内容")
	cron.NewTwLolLuckDraw().Run()
	//添加定时任务 cron表达式文档 https://pkg.go.dev/github.com/robfig/cron
	c := crontab.New(crontab.WithSeconds())
	cronList := []*CronList{
		{"0 * * * * *", cron.NewTwLolLuckDraw().Run},
	}
	for _, v := range cronList {
		_, err := c.AddFunc(v.spec, v.run)
		if err != nil {
			panic(err)
		}
		c.Start()
	}
}
