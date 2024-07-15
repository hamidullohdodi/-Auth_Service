package logger

import (
	"path/filepath"

	"go.uber.org/zap"
)

func developmentConfig(file string) zap.Config {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.OutputPaths = []string{"stdout", file}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	return zapConfig
}

func New(level, enviroment, file_name string) (*zap.Logger, error) {
	file := filepath.Join("./" + file_name)

	zapConfig := developmentConfig(file)
	switch level {
	case "debug":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		zapConfig.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	return zapConfig.Build()
}
