package util

import (
	"encoding/json"
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
	jsonByt, _ := json.Marshal(content)
	l.drive.Info(msg, zap.String("content", string(jsonByt)))
}

func (l *logger) Error(msg string, content Map) {
	jsonByt, _ := json.Marshal(content)
	l.drive.Error(msg, zap.String("content", string(jsonByt)))
}
