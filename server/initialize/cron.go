package initialize

import (
	"app/cron"
	crontab "github.com/robfig/cron/v3"
)

type CronList struct {
	spec string
	run  func()
}

func Cron() {
	//添加定时任务 cron表达式文档 https://pkg.go.dev/github.com/robfig/cron
	c := crontab.New(crontab.WithSeconds())
	cronList := []*CronList{
		{"0 0 */8 * * *", cron.NewTwLolLuckDraw().Run},
	}
	for _, v := range cronList {
		_, err := c.AddFunc(v.spec, v.run)
		if err != nil {
			panic(err)
		}
		c.Start()
	}
}
