package initialize

import (
	"app/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//获取redis连接字符串
func getRedisConnString() string {
	return fmt.Sprintf("%s:%d",
		global.Config.Redis.Host,
		global.Config.Redis.Port)
}

func Redis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:         getRedisConnString(),
		Password:     global.Config.Redis.Password,
		DB:           global.Config.Redis.Database,
		PoolSize:     global.Config.Redis.MaxOpenConn, //最大连接数
		MinIdleConns: global.Config.Redis.MaxIdleConn, //最大空闲连接数
	})
	_, err := global.Redis.Ping(context.Background()).Result()
	if err != nil {
		panic("redis conn error: " + err.Error())
	}
}
