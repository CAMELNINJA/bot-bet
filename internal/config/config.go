package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string           `yaml:"env" env:"ENV" env-default:"development"`
	Telegram   Telegram         `yaml:"telegram"`
	HttpServer HttpServerConfig `yaml:"http"`
	Postgres   PostgresConfig   `yaml:"postgres"`
}

type Telegram struct {
	BotToken  string `yaml:"bot_token" env:"TELEGRAM_BOT_TOKEN"`
	WebAppUrl string `yaml:"web_app_url" env:"TELEGRAM_WEB_APP_URL"`
	Timeout   int    `yaml:"timeout" env:"TELEGRAM_TIMEOUT" env-default:"60"`
	Admins    []int  `yaml:"admins" env:"TELEGRAM_ADMINS"`
}

type HttpServerConfig struct {
	Address     string        `yaml:"addres" env:"HTTP_SERVER_ADDRESS" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_SERVER_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idletimeot" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"60s"`
}

type PostgresConfig struct {
	Host          string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port          int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	User          string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	Password      string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
	Name          string `yaml:"dbname" env:"POSTGRES_DBNAME" env-default:"postgres"`
	MigrateSource string `yaml:"migrate_source" env:"POSTGRES_MIGRATE_SOURCE" env-default:"file://migrations"`
}

func MustLoad(configPath string) (*Config, error) {
	var config Config
	if configPath == "" {
		slog.Info("configPath not found")
		if err := cleanenv.ReadEnv(&config); err != nil {
			return nil, err
		}
	} else {
		slog.Info("configPath ", slog.String("path", configPath))
		if _, err := os.Stat(configPath); err != nil {
			slog.Error("err cheking file", err)
			return nil, err
		}
		if err := cleanenv.ReadConfig(configPath, &config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}

func (c *Config) ConnString() string {
	uri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.User, c.Postgres.Name, c.Postgres.Password)

	return uri
}

func (c *Config) Address() (string, error) {
	if len(strings.Split(c.HttpServer.Address, ":")) != 2 {
		return "", fmt.Errorf("invalid address %s", c.HttpServer.Address)
	}
	return c.HttpServer.Address, nil
}
