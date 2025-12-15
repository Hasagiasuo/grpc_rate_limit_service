package logger

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger(env string) *Logger {
	log := setupLogger(env)
	return &Logger{log}
}

func (l *Logger) Info(op, message string) {
	l.log.Info(fmt.Sprintf("%s => %s;", op, message))
}

func (l *Logger) Warn(op, message string) {
	l.log.Warn(fmt.Sprintf("%s => %s;", op, message))
}

func (l *Logger) Error(op, message string) {
	l.log.Error(fmt.Sprintf("%s => %s;", op, message))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
