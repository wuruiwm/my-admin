package initialize

import (
	"app/task"
)

func Task() {
	task.AdminConfig()
	go task.Youtube()
	go task.M3u8()
}
