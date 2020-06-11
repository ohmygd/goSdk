package main

import (
	logger "go.uber.org/zap"
	"time"

	rotatelogs "github.com/eleven-gif/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	infoLogPath   = "E:\\mc.log-%Y-%m-%dT%H.info"
	warnLogPath   = "E:\\mc.log-%Y-%m-%dT%H.warn"
	publicLogPath = "E:\\public.log"
	openDebug     = false
)

//Initlog 初始化log.debug和info输出一个文件，warn，error输出一个文件
func InitLog() {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	loginfo, _ := rotatelogs.New(
		infoLogPath,
		rotatelogs.WithLinkName("mc.log.info"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	logwarn, _ := rotatelogs.New(
		warnLogPath,
		rotatelogs.WithLinkName("mc.log.warn"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	infoLog := zapcore.AddSync(loginfo)
	warnLog := zapcore.AddSync(logwarn)
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			infoLog,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				if openDebug {
					return lvl <= zapcore.InfoLevel
				}
				return lvl == zapcore.InfoLevel
			},
			),
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			warnLog,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.WarnLevel
			},
			),
		),
	)

	logger := zap.New(core)
	logger = logger.WithOptions(zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func main() {
	InitLog()
	logger.L().Info("hello")
}
