package util

import (
	"app/global"
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
	l.logger.Info(msg, zap.Any("content", content))
}

func (l *Logger) Error(msg string, content Map) {
	l.logger.Error(msg, zap.Any("content", content))
}
