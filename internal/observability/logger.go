package observability

import (
	"go.uber.org/zap"
)

// Logger defines logging interface for dependency injection.
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
}

type zapLogger struct {
	logger *zap.Logger
}

// NewLogger creates a production zap logger.
func NewLogger() (Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &zapLogger{logger: l}, nil
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *zapLogger) Sync() error {
	return l.logger.Sync()
}