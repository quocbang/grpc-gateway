package logging

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
)

type LogLevel struct {
	level logger.LogLevel
}

func NewGormLogger() logger.Interface {
	return &LogLevel{
		level: logger.Info,
	}
}

// LogMode log mode
func (l *LogLevel) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.level = level
	return &newlogger
}

// Error print error messages
func (l LogLevel) Error(ctx context.Context, msg string, i ...interface{}) {
	if l.level >= logger.Error {
		interceptors.GetLoggerFormContext(ctx).Error(fmt.Sprintf(msg, i...), zap.String("caller", utils.FileWithLineNum()))
	}
}

// Info print info
func (l LogLevel) Info(ctx context.Context, msg string, i ...interface{}) {
	if l.level >= logger.Info {
		interceptors.GetLoggerFormContext(ctx).Info(fmt.Sprintf(msg, i...), zap.String("caller", utils.FileWithLineNum()))
	}
}

// Warn print warn messages
func (l LogLevel) Warn(ctx context.Context, msg string, i ...interface{}) {
	if l.level >= logger.Warn {
		interceptors.GetLoggerFormContext(ctx).Warn(fmt.Sprintf(msg, i...), zap.String("caller", utils.FileWithLineNum()))
	}
}

// Trace print sql message
func (l LogLevel) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := []zap.Field{
		zap.String("caller", utils.FileWithLineNum()),
		zap.Duration("elapsed_time", elapsed),
		zap.String("sql", sql),
		zap.Int64("rowsAffected", rows),
	}
	if err != nil && l.level >= logger.Error {
		fields = append(fields, zap.Error(err))
		interceptors.GetLoggerFormContext(ctx).Error("tracing SQL..", fields...)
	} else {
		interceptors.GetLoggerFormContext(ctx).Info("tracing SQL", fields...)
	}
}

type RedisLogger struct{}

type Logging interface {
	Printf(ctx context.Context, format string, v ...interface{})
}

func NewRedisLogger() Logging {
	return &RedisLogger{}
}

// Printf is print redis log.
func (rl *RedisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	interceptors.GetLoggerFormContext(ctx).Info("Tracing Redis SQL..", zap.String("sql", fmt.Sprintf(format, v...)))
}
