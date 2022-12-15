package crontab

import (
	"app/util"
	"fmt"
)

func Minute() {
	fmt.Printf("分钟级定时任务：%s\n", util.Date())
}
