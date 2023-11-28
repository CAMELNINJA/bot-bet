package bot

import (
	"log/slog"
	"os"

	"github.com/CAMELNINJA/bot-bet.git/internal/config"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "prod"
)

func StartBot(configPath string) error {
	cfg, err := config.MustLoad(configPath)
	if err != nil {
		return err
	}
	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.HttpServer.Address))
	log.Debug("logger debug mode enabled")

	return nil
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
