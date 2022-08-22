package middleware

import (
	"app/global"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"time"
)

func Logger(c *gin.Context) {
	//获取请求开始时间
	start := time.Now()
	//获取body内容 因为body只能读一次 所以再写回body
	bodyByt, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByt))
	body := make(map[string]interface{}, 0)
	_ = json.Unmarshal(bodyByt, &body)

	c.Next()

	//获取请求结束时间
	cost := time.Since(start)
	fields := []zapcore.Field{
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("query", c.Request.URL.RawQuery),
		zap.Any("body", body),
		zap.String("ip", c.ClientIP()),
		zap.Duration("cost", cost),
	}

	//recover捕获到panic之后 将捕获到的error写入panic字段
	var isError bool
	panicErr, ok := c.Get("panic")
	if ok {
		isError = true
		fields = append(fields, zap.Any("panic", panicErr))
	}
	//controller返回的错误 写入到error字段
	err, ok := c.Get("error")
	if ok {
		isError = true
		fields = append(fields, zap.Any("error", err))
	}

	//写入日志
	if isError {
		global.Logger.Error(c.Request.URL.Path, fields...)
	} else {
		global.Logger.Info(c.Request.URL.Path, fields...)
	}
}
