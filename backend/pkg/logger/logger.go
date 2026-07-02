package logger

import (
	"go.uber.org/zap"
)

// Logger wraps the Zap logger.
type Logger struct {
	*zap.Logger
}

// New creates a new logger instance.
func New(environment string) (*Logger, error) {
	var (
		zapLogger *zap.Logger
		err       error
	)

	switch environment {
	case "production":
		zapLogger, err = zap.NewProduction()
	default:
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: zapLogger,
	}, nil
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() error {
	if l == nil || l.Logger == nil {
		return nil
	}

	return l.Logger.Sync()
}