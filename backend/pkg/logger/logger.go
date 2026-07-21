package logger

import "go.uber.org/zap"

// Logger wraps a Zap logger.
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

// Named returns a child logger with a component name.
func (l *Logger) Named(name string) *Logger {
	if l == nil || l.Logger == nil {
		return l
	}

	return &Logger{
		Logger: l.Logger.Named(name),
	}
}

// With returns a child logger with structured fields.
func (l *Logger) With(fields ...zap.Field) *Logger {
	if l == nil || l.Logger == nil {
		return l
	}

	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if l != nil && l.Logger != nil {
		l.Logger.Debug(msg, fields...)
	}
}

// Info logs an informational message.
func (l *Logger) Info(msg string, fields ...zap.Field) {
	if l != nil && l.Logger != nil {
		l.Logger.Info(msg, fields...)
	}
}

// Warn logs a warning.
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if l != nil && l.Logger != nil {
		l.Logger.Warn(msg, fields...)
	}
}

// Error logs an error.
func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l != nil && l.Logger != nil {
		l.Logger.Error(msg, fields...)
	}
}

// Fatal logs a fatal error and exits.
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	if l != nil && l.Logger != nil {
		l.Logger.Fatal(msg, fields...)
	}
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() error {
	if l == nil || l.Logger == nil {
		return nil
	}

	return l.Logger.Sync()
}
