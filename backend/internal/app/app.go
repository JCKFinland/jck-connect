package app

import (
	"fmt"
	"net/http"

	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/internal/router"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
	"github.com/JCKFinland/jck-connect/backend/pkg/logger"
)

type App struct {
	Config *config.Config
	Logger *logger.Logger
	DB     *database.Database
	Server *http.Server
}

func New() (*App, error) {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log, err := logger.New(cfg.AppEnv)
	if err != nil {
		return nil, fmt.Errorf("initialize logger: %w", err)
	}

	// Connect database
	db, err := database.New(cfg)
	if err != nil {
		log.Error("database connection failed")
		return nil, fmt.Errorf("connect database: %w", err)
	}

	// Build router
	r := router.New()

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	return &App{
		Config: cfg,
		Logger: log,
		DB:     db,
		Server: server,
	}, nil
}

func (a *App) Run() error {
	a.Logger.Info("Starting jck-connect server",
		logger.String("port", a.Config.AppPort),
	)

	return a.Server.ListenAndServe()
}

func (a *App) Shutdown() error {
	a.Logger.Info("Shutting down application")

	if a.DB != nil {
		a.DB.Close()
	}

	return a.Logger.Sync()
}