package pkg

import (
	"fmt"
	"gin_template/internal/app/config"
	"os"
	"time"

	"gin_template/internal/app/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(config.C.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.C.Zap.Director)
		_ = os.Mkdir(config.C.Zap.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", config.C.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", config.C.Zap.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", config.C.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", config.C.Zap.Director), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if config.C.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (cf zapcore.EncoderConfig) {
	cf = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.C.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case config.C.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		cf.EncodeLevel = zapcore.LowercaseLevelEncoder
	case config.C.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		cf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case config.C.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		cf.EncodeLevel = zapcore.CapitalLevelEncoder
	case config.C.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		cf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		cf.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return cf
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if config.C.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := utils.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(config.C.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
