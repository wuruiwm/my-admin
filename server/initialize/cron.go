package initialize

import "github.com/robfig/cron/v3"

type CronList struct {
	spec string
	run  func()
}

func Cron() {
	//添加定时任务 cron表达式文档 https://pkg.go.dev/github.com/robfig/cron
	c := cron.New(cron.WithSeconds())
	cronList := []*CronList{
		//{"* * * * * *", task.SecondCron}, //秒  级 定时任务
		//{"0 * * * * *", task.MinuteCron}, //分钟级 定时任务
	}
	for _, v := range cronList {
		_, err := c.AddFunc(v.spec, v.run)
		if err != nil {
			panic(err)
		}
		c.Start()
	}
}
