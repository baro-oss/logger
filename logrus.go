package logger

import (
	"context"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type LrLogger struct {
	logger *logrus.Logger
}

func NewLrLogger() *LrLogger {
	logger := logrus.New()
	logger.SetOutput(os.Stderr)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyFunc: "caller",
		},
	})
	logger.ReportCaller = true

	return &LrLogger{
		logger: logger,
	}
}

func (l *LrLogger) Sync() {
	return
}

func (l *LrLogger) Info(msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg)))
	l.logger.Info(logrus.WithFields(logFields))
}

func (l *LrLogger) Warn(msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg)))
	l.logger.Warn(logrus.WithFields(logFields))
}

func (l *LrLogger) Err(msg string, fields ...Field) {
	logFields := convertLrFields(fields)
	l.logger.WithError(errors.New(msg)).Error(logrus.WithFields(logFields))
}

func (l *LrLogger) Fatal(msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg)))
	l.logger.Fatal(logrus.WithFields(logFields))
}

func (l *LrLogger) Debug(msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg)))
	l.logger.Debug(logrus.WithFields(logFields))
}

func (l *LrLogger) Trace(msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg)))
	l.logger.Trace(logrus.WithFields(logFields))
}

func (l *LrLogger) InfoWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg), WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).Info(logrus.WithFields(logFields))
}

func (l *LrLogger) WarnWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg), WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).Warn(logrus.WithFields(logFields))
}

func (l *LrLogger) ErrWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).WithError(errors.New(msg)).Info(logrus.WithFields(logFields))
}

func (l *LrLogger) FatalWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg), WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).Fatal(logrus.WithFields(logFields))
}

func (l *LrLogger) DebugWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg), WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).Debug(logrus.WithFields(logFields))
}

func (l *LrLogger) TraceWithCtx(ctx context.Context, msg string, fields ...Field) {
	logFields := convertLrFields(append(fields, WithField(LogFieldMsg, msg), WithField(TraceIdKey, ctx.Value(TraceIdKey))))
	l.logger.WithContext(ctx).Trace(logrus.WithFields(logFields))
}

func convertLrFields(fields []Field) logrus.Fields {
	var logFields logrus.Fields
	for _, field := range fields {
		logFields[field.key] = field.value
	}
	logFields[LogFieldDriver] = LogDriverLogrus
	return logFields
}
