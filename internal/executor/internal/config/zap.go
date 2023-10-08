package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DefaultZap() zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig = zap.NewProductionEncoderConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return config
}
