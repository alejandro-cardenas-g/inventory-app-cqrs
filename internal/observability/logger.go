package observability

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Service string
	Env     string
	Level   string
}

func New(cfg Config) (*zap.Logger, error) {
	logger, err := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	logger.With(zap.String("service", cfg.Service), zap.String("env", cfg.Env))

	if err != nil {
		return nil, err
	}

	return logger, nil
}