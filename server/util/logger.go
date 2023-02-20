package util

import (
	"app/global"
	"encoding/json"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logger ...*zap.Logger) *Logger {
	if len(logger) == 0 {
		return &Logger{
			logger: global.Logger,
		}
	} else if len(logger) == 1 {
		return &Logger{
			logger: logger[0],
		}
	} else {
		panic("logger param num error")
	}
}

func (l *Logger) Info(msg string, content Map) {
	jsonByt, _ := json.Marshal(content)
	l.logger.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *Logger) Error(msg string, content Map) {
	jsonByt, _ := json.Marshal(content)
	l.logger.Error(msg, zap.String("content", string(jsonByt)))
}
