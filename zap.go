package logger

import (
	"context"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger   *zap.Logger
	isGlobal bool
}

func NewZapLogger(withGlobal bool) *ZapLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	config.EncoderConfig.TimeKey = "@timestamp"
	logger, _ := config.Build()

	if withGlobal {
		zap.ReplaceGlobals(logger)
	}

	return &ZapLogger{
		logger:   logger,
		isGlobal: withGlobal,
	}
}

func (l *ZapLogger) Sync() {
	if l.isGlobal {
		zap.L().Sync()
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Sync()
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Info(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Info(msg, convertZapField(fields)...)
}

func (l *ZapLogger) Warn(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Warn(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Warn(msg, convertZapField(fields)...)
}

func (l *ZapLogger) Err(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Error(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Error(msg, convertZapField(fields)...)
}

func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Fatal(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Fatal(msg, convertZapField(fields)...)
}

func (l *ZapLogger) Debug(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Debug(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Debug(msg, convertZapField(fields)...)
}

func (l *ZapLogger) Trace(msg string, fields ...Field) {
	if l.isGlobal {
		zap.L().Info(msg, convertZapField(fields)...)
		return
	}

	if l.logger == nil {
		return
	}

	l.logger.Info(msg, convertZapField(fields)...)
}

func (l *ZapLogger) InfoWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Info(msg, fields...)
}

func (l *ZapLogger) WarnWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Warn(msg, fields...)
}

func (l *ZapLogger) ErrWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Err(msg, fields...)
}

func (l *ZapLogger) FatalWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Fatal(msg, fields...)
}

func (l *ZapLogger) DebugWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Debug(msg, fields...)
}

func (l *ZapLogger) TraceWithCtx(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, WithField(TraceIdKey, ctx.Value(TraceIdKey)))
	l.Trace(msg, fields...)
}

func convertZapField(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields)+1)

	for i, field := range fields {
		switch reflect.TypeOf(field.value).Kind() {
		case reflect.Bool:
			zapFields[i] = zap.Bool(field.key, field.value.(bool))
		case reflect.Int:
			zapFields[i] = zap.Int(field.key, field.value.(int))
		case reflect.Int64:
			zapFields[i] = zap.Int64(field.key, field.value.(int64))
		case reflect.Int32:
			zapFields[i] = zap.Int32(field.key, field.value.(int32))
		case reflect.String:
			zapFields[i] = zap.String(field.key, field.value.(string))
		case reflect.Pointer:
			val := reflect.ValueOf(field.value)
			if !val.IsZero() {
				zapFields[i] = zap.Any(field.key, val.Elem().Interface())
				continue
			}
			zapFields[i] = zap.Any(field.key, field.value)
		default:
			zapFields[i] = zap.Any(field.key, field.value)
		}
	}

	zapFields[len(zapFields)-1] = zap.String(LogFieldDriver, LogDriverZap)

	return zapFields
}
