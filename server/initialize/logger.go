package initialize

import (
	"app/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 初始化Log
func Logger() {
	writeSyncer := loggerLogWriter()
	encoder := loggerEncoder()
	var log = new(zapcore.Level)
	err := log.UnmarshalText([]byte("info"))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSyncer, log)
	global.Logger = zap.New(core, zap.AddCaller())
}

func loggerEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "line"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func loggerLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/server.log", // 保存的文件位置
		MaxSize:    100,                // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 100,                // 保留归档日志文件的最大个数
		MaxAge:     30,                 // 保留归档日志文件的最大天数
		Compress:   false,              // 是否压缩归档日志文件
		LocalTime:  true,               // 采用本地时区
	}
	return zapcore.AddSync(lumberJackLogger)
}
