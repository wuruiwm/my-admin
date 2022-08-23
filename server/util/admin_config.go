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
	cacheKey := "admin_config:" + arr[0] + ":" + arr[1]
	val, err := global.Redis.Get(context.Background(), cacheKey).Result()
	if err != nil {
		global.Db.Table("admin_config").
			Where("group", arr[0]).
			Where("key", arr[1]).
			Select("value").
			Scan(&val)
		_ = global.Redis.Set(context.Background(), cacheKey, val, 86400).Err()
	}
	return val
}
