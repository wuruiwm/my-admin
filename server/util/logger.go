package util

import (
	"app/global"
	"github.com/bytedance/sonic/encoder"
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
	jsonByt, _ := encoder.Encode(content, encoder.SortMapKeys)
	global.Logger.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *Logger) Error(msg string, content Map) {
	jsonByt, _ := encoder.Encode(content, encoder.SortMapKeys)
	global.Logger.Error(msg, zap.String("content", string(jsonByt)))
}
