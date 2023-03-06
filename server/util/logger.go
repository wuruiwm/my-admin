package util

import (
	"app/global"
	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

var logger *Logger

type Logger struct {
}

func NewLogger() *Logger {
	if logger == nil {
		logger = &Logger{}
	}
	return logger
}

func (l *Logger) Info(msg string, content Map) {
	jsonByt, _ := sonic.Marshal(content)
	global.Logger.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *Logger) Error(msg string, content Map) {
	jsonByt, _ := sonic.Marshal(content)
	global.Logger.Error(msg, zap.String("content", string(jsonByt)))
}
