package initialize

import "app/task"

func Task() {
	go task.AdminConfig()
}
