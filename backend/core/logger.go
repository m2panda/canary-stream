package core

import (
	"log/slog"
	"os"
)

func InitLoggerService() {
	var handler slog.Handler

	environment := os.Getenv("ENV_MODE")

	if environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
