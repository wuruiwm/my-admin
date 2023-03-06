package middleware

import (
	"app/util"
	"bytes"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func Logger(c *gin.Context) {
	//获取请求开始时间
	t := time.Now()
	//获取body内容 因为body只能读一次 所以再写回body
	bodyByt, _ := c.GetRawData()
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByt))
	body := make(map[string]interface{}, 0)
	_ = sonic.Unmarshal(bodyByt, &body)

	c.Next()

	//获取请求结束时间
	logContent := util.Map{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"query":  c.Request.URL.RawQuery,
		"body":   body,
		"ip":     c.ClientIP(),
		"cost":   util.TimeSince(t),
	}

	//recover捕获到panic之后 将捕获到的error写入panic字段
	var isError bool
	panicErr, ok := c.Get("panic")
	if ok {
		isError = true
		logContent["panic"] = panicErr
	}
	//controller返回的错误 写入到error字段
	err := c.GetString("error")
	if err != "" {
		isError = true
		logContent["error"] = err
	}

	//写入日志
	msg := "api"
	if isError {
		util.NewLogger().Error(msg, logContent)
	} else {
		util.NewLogger().Info(msg, logContent)
	}
}
