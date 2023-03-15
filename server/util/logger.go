package util

import (
	"github.com/bytedance/sonic/encoder"
	"go.uber.org/zap"
)

var Logger *logger

type logger struct {
	drive *zap.Logger
}

func InitLogger(drive *zap.Logger) {
	Logger = &logger{
		drive: drive,
	}
}

func (l *logger) Info(msg string, content Map) {
	jsonByt, _ := encoder.Encode(content, encoder.SortMapKeys)
	l.drive.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *logger) Error(msg string, content Map) {
	jsonByt, _ := encoder.Encode(content, encoder.SortMapKeys)
	l.drive.Error(msg, zap.String("content", string(jsonByt)))
}
