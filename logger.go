package logger

import (
	"context"
)

const (
	TraceIdKey     = "X-Trace-ID"
	LogFieldMsg    = "message"
	LogFieldDriver = "log-driver"

	LogDriverZap     = "zap"
	LogDriverLogrus  = "logrus"
	LogDriverZeroLog = "zerolog"
)

type LogDriver string

type Logger interface {
	Sync()

	Info(msg string, Fields ...Field)
	Warn(msg string, Fields ...Field)
	Err(msg string, Fields ...Field)
	Fatal(msg string, Fields ...Field)
	Debug(msg string, Fields ...Field)
	Trace(msg string, Fields ...Field)

	InfoWithCtx(ctx context.Context, msg string, Fields ...Field)
	WarnWithCtx(ctx context.Context, msg string, Fields ...Field)
	ErrWithCtx(ctx context.Context, msg string, Fields ...Field)
	FatalWithCtx(ctx context.Context, msg string, Fields ...Field)
	DebugWithCtx(ctx context.Context, msg string, Fields ...Field)
	TraceWithCtx(ctx context.Context, msg string, Fields ...Field)
}

// NewLogger return a concrete of Logger interface rely on log driver provided in argument.
// Function NewLogger supports log drivers include LogDriverZap, LogDriverLogrus, LogDriverZeroLog
// default logger instance that created if an invalid log driver provided is ZapLogger
func NewLogger(logDriver LogDriver, withGlobal bool) Logger {
	switch logDriver {
	case LogDriverZap:
		return NewZapLogger(withGlobal)
	case LogDriverLogrus:
		return NewLrLogger()
	case LogDriverZeroLog:
		return NewZeroLogger()
	default:
		return NewZapLogger(withGlobal)
	}
}

type Field struct {
	key   string
	value any
}

func WithField(key string, value any) Field {
	return Field{
		key:   key,
		value: value,
	}
}
