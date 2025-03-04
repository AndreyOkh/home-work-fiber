package logger

import (
	"log/slog"
	"news/config"
	"os"
)

func NewLogger(conf *config.LogConfig) *slog.Logger {
	lvl := new(slog.LevelVar)

	var logger *slog.Logger
	switch conf.Format {
	case "text":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	default:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}))
	}

	switch conf.Level {
	case "debug":
		lvl.Set(slog.LevelDebug)
	case "info":
		lvl.Set(slog.LevelInfo)
	case "warn":
		lvl.Set(slog.LevelWarn)
	case "error":
		lvl.Set(slog.LevelError)
	default:
		slog.Info("Undefined log level, defaulting to info")
		lvl.Set(slog.LevelInfo)
	}

	return logger
}
