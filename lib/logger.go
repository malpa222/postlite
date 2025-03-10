package lib

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(debug bool) Logger {
	return Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
	}
}

func (l Logger) Debug(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args)
	l.logger.Debug(msg)
}

func (l Logger) Info(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args)
	l.logger.Debug(msg)
}

func (l Logger) Warn(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args)
	l.logger.Warn(msg)
}

func (l Logger) Error(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args)
	l.logger.Error(msg)
}

func (l Logger) Fatal(msg string, args ...any) {
	msg = fmt.Sprintf(msg, args)
	msg = fmt.Sprintf("Unrecoverable error: %s", msg)

	panic(msg)
}
