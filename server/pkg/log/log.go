package log

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func Close() error {
	return logger.Sync()
}

func Info(msg string) {
	logger.Info(msg)
}

func Warm(msg string) {
	logger.Warn(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Fatal(msg string) {
	logger.Fatal(msg)
}
