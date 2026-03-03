package observability

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

// InitLogger initializes the logger with the specified configuration.
func InitLogger() error {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}

// Info logs an informational message.
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn logs a warning message.
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error logs an error message.
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits the application.
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries.
func Sync() {
	_ = logger.Sync()
}