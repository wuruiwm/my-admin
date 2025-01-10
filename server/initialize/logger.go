package initialize

import (
	"app/util"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
	"strings"
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
	util.InitLogger(zap.New(core, zap.AddCaller()))
}

func loggerEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "line"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = shortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func loggerLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/server.log", // 保存的文件位置
		MaxSize:    100,                // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,                 // 保留归档日志文件的最大个数
		MaxAge:     90,                 // 保留归档日志文件的最大天数
		Compress:   false,              // 是否压缩归档日志文件
		LocalTime:  true,               // 采用本地时区
	}
	return zapcore.AddSync(lumberJackLogger)
}

func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	_, logFile, _, ok := runtime.Caller(0)
	if !ok {
		enc.AppendString("")
		return
	}
	baseDir := strings.Replace(logFile, "initialize/logger.go", "", 1)
	for i := 2; i < 999; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && strings.HasPrefix(file, baseDir) && !strings.HasSuffix(file, "/util/logger.go") && !strings.HasSuffix(file, "/initialize/db.go") {
			f := strings.Replace(file, baseDir, "", 1)
			enc.AppendString(f + ":" + cast.ToString(line))
			return
		}
	}
	enc.AppendString("")
}
