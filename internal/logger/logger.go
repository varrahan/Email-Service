package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() {
	cfg := zap.NewProductionConfig()

	if os.Getenv("ENV") == "dev" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Build logger
	logger, err := cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}

	Log = logger
}

func Sync() {
	_ = Log.Sync() // nolint:errcheck
}