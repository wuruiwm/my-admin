package util

import (
	"app/global"
	"context"
	"strings"
)

func AdminConfig(key string) string {
	arr := strings.Split(key, ".")
	if len(arr) != 2 {
		return ""
	}
	val, err := global.Redis.Get(context.Background(), "admin_config:"+arr[0]+":"+arr[1]).Result()
	if err != nil {
		global.Db.Table("admin_config").
			Where("group", arr[0]).
			Where("key", arr[1]).
			Select("value").
			Scan(&val)
	}
	return val
}
