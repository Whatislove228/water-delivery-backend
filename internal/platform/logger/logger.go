package logger

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	switch strings.ToLower(level) {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return nil, fmt.Errorf("unsupported log level: %s", level)
	}

	return cfg.Build()
}
