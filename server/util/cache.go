package util

import (
	"app/global"
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"time"
)

// CacheSet 设置缓存
func CacheSet(key string, data interface{}, expire int64) error {
	jsonString, err := sonic.Marshal(data)
	if err != nil {
		return errors.New("序列化为json字符串失败 error:" + err.Error())
	}
	err = global.Redis.Set(context.Background(), key, jsonString, time.Duration(expire)*time.Second).Err()
	if err != nil {
		return errors.New("设置缓存失败 error:" + err.Error())
	}
	return nil
}

// CacheGet 获取缓存
func CacheGet(key string, data interface{}) error {
	jsonString, err := global.Redis.Get(context.Background(), key).Result()
	if err != nil {
		return errors.New("获取缓存失败 error:" + err.Error())
	}
	err = sonic.Unmarshal([]byte(jsonString), data)
	if err != nil {
		return errors.New("json字符串反序列化失败 error:" + err.Error())
	}
	return nil
}
