package util

import (
	"github.com/bytedance/sonic"
	"go.uber.org/zap"
)

var Logger *logger

var json = sonic.Config{
	SortMapKeys: true,
}.Froze()

type logger struct {
	drive *zap.Logger
}

func InitLogger(drive *zap.Logger) {
	Logger = &logger{
		drive: drive,
	}
}

func (l *logger) Info(msg string, content Map) {
	jsonByt, _ := json.Marshal(content)
	l.drive.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *logger) Error(msg string, content Map) {
	jsonByt, _ := json.Marshal(content)
	l.drive.Error(msg, zap.String("content", string(jsonByt)))
}
