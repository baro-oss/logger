package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	logger *zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "@timestamp"

	logger := zerolog.New(os.Stderr)

	return &ZeroLogger{
		logger: &logger,
	}
}

func (l *ZeroLogger) Sync() {
	return
}

func (l *ZeroLogger) Info(msg string, fields ...Field) {
	l.logger.Info().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) Err(msg string, fields ...Field) {
	l.logger.Error().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) Trace(msg string, fields ...Field) {
	l.logger.Trace().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) InfoWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Info().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) WarnWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Warn().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) ErrWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Error().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) FatalWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Fatal().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) DebugWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Debug().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func (l *ZeroLogger) TraceWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	zerolog.Ctx(l.logger.WithContext(ctx)).Trace().Fields(convertZeroLogFields(fields)).Msg(msg)
}

func convertZeroLogFields(fields []Field) map[string]any {
	logFields := make(map[string]any)
	for _, field := range fields {
		logFields[field.key] = field.value
	}
	logFields[LogFieldDriver] = LogDriverZeroLog
	return logFields
}
