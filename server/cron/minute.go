package cron

import (
	"fmt"
	"time"
)

func MinuteCron() {
	fmt.Printf("分钟级定时任务：%s\n", time.Now().Format("2006-01-02 03:04:05"))
}
