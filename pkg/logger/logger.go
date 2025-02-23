package logger

import (
	"log/slog"
	"news/config"
	"os"
)

func NewLogger(conf *config.LogConfig) *slog.Logger {
	level := 0
	switch conf.Level {
	case "debug":
		level = -4
	case "info":
		level = 0
	case "warn":
		level = 4
	case "error":
		level = 8
	default:
		slog.Info("Undefined log level, defaulting to info")
		level = 0
	}
	slog.SetLogLoggerLevel(slog.Level(level))

	switch conf.Format {
	case "text":
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
}
