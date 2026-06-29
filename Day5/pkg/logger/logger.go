package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.Logger

func Init(env string) {
	var core zapcore.Core

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace", //完整的调用栈信息
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if env == "production" {
		//判断是不是生成环境
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)
	} else {
		//开发环境
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)
	}
	global = zap.New(core, zap.AddCaller())
}

func Info(msg string, fields ...zap.Field)  { global.Info(msg, fields...) }
func Error(msg string, fields ...zap.Field) { global.Error(msg, fields...) }
func Debug(msg string, fields ...zap.Field) { global.Debug(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { global.Warn(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { global.Fatal(msg, fields...) }
func With(fields ...zap.Field) *zap.Logger  { return global.With(fields...) }
