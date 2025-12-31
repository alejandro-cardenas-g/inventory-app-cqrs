package observability

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Service string
	Env     string
	Level   string
}

func New(cfg Config) (*zap.Logger, error) {
	level := parseLevel(cfg.Level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(zapcore.Lock(zapcore.AddSync(zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(zapcore.AddSync(nil)),
		)))),
		level,
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	).With(
		zap.String("service", cfg.Service),
		zap.String("env", cfg.Env),
	)
	
	return logger, nil
}

func parseLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}