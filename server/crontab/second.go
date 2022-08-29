package crontab

import (
	"fmt"
	"time"
)

func Second() {
	fmt.Printf("秒级定时任务：%s\n", time.Now().Format("2006-01-02 03:04:05"))
}
