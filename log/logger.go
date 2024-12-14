package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = zap.NewNop()

func InitLogger(cfg Config) error {
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "log",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return err
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	Logger.Sugar().Fatalf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	Logger.Sugar().Errorf(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	Logger.Sugar().Infof(template, args...)
}
