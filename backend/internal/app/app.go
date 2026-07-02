package app

import (
	"github.com/JCKFinland/jck-connect/backend/internal/bootstrap"
)

type App struct {
	Config *bootstrap.Config
	Logger *bootstrap.Logger
	DB     *bootstrap.Database
}

func New() (*App, error) {
	cfg := bootstrap.LoadConfig()

	logger := bootstrap.NewLogger(cfg.AppEnv)

	db, err := bootstrap.NewDatabase(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Logger: logger,
		DB:     db,
	}, nil
}