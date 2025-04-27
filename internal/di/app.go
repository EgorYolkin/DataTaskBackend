package di

import (
	"DataTask/internal/config"
	"database/sql"
)

type App struct {
	Config   *config.Config
	Database *sql.DB
}

func InitializeApp(cfg *config.Config) (*App, error) {
	return &App{
		Config: cfg,
	}, nil
}
