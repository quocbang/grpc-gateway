package logging

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AsynqLog struct{}

func NewAsynqLog() *AsynqLog {
	return &AsynqLog{}
}

func (a *AsynqLog) Print(level zerolog.Level, args ...interface{}) {
	log.WithLevel(level).Msg(fmt.Sprint(args...))
}

// Debug logs a message at Debug level.
func (a *AsynqLog) Debug(args ...interface{}) {
	a.Print(zerolog.DebugLevel, args...)
}

// Info logs a message at Info level.
func (a *AsynqLog) Info(args ...interface{}) {
	a.Print(zerolog.InfoLevel, args...)
}

// Warn logs a message at Warning level.
func (a *AsynqLog) Warn(args ...interface{}) {
	a.Print(zerolog.WarnLevel, args...)
}

// Error logs a message at Error level.
func (a *AsynqLog) Error(args ...interface{}) {
	a.Print(zerolog.ErrorLevel, args...)
}

// Fatal logs a message at Fatal level
// and process will exit with status set to 1.
func (a *AsynqLog) Fatal(args ...interface{}) {
	a.Print(zerolog.FatalLevel, args...)
}
