package initialize

import "app/task"

func Task() {
	task.AdminConfig()
	go task.TwLolLuckDraw()
}
