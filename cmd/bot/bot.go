package bot

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/CAMELNINJA/bot-bet.git/internal/config"
	httpServer "github.com/CAMELNINJA/bot-bet.git/internal/http"
	postgresrepo "github.com/CAMELNINJA/bot-bet.git/internal/repositories/postgres_repo"
	"github.com/CAMELNINJA/bot-bet.git/internal/telegram"
	"github.com/CAMELNINJA/bot-bet.git/internal/usecase"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
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

	db, err := setUpPostgres(log, cfg)
	if err != nil {
		return err
	}

	repo := postgresrepo.NewPostrgesRepo(db, nil, log)

	userUsecase := usecase.NewUserUsecase(log, repo)
	gameUsecase := usecase.NewGame(repo, log)
	gameHandler := httpServer.NewGameHandler(gameUsecase, log)

	bot := telegram.NewAdapter(log, userUsecase, cfg)
	r := httpServer.NewRouter(log, gameHandler)
	addr, err := cfg.Address()
	if err != nil {
		return err
	}
	errCh := make(chan error, 1)
	go func() {
		if err := bot.Listener(); err != nil {
			errCh <- err
		}
	}()

	go func() {
		err = http.ListenAndServe(addr, r)
		if err != nil {
			log.Error("Error starting server", slog.String("error", err.Error()))
			errCh <- err
		}
	}()

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

func setUpPostgres(log *slog.Logger, cfg *config.Config) (*sqlx.Db, error) {
	log.Info("logerr add  successfully!")
	// Connect to the database
	db, err := sqlx.Connect("pgx", cfg.ConnString())
	if err != nil {
		log.Error("Error connecting to database", slog.String("error", err.Error()))
		return nil, err
	}

	// Migrations block
	err = goose.Up(db.DB, cfg.Postgres.MigrateSource)
	if err != nil {
		log.Error("error migration", slog.String("err", err.Error()))
		return nil, err
	}
	log.Info("migrations successfully!")
	return db, nil
}
