package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

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

// New initializes the application.
func New() (*App, error) {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log, err := logger.New(cfg.AppEnv)
	if err != nil {
		return nil, fmt.Errorf("initialize logger: %w", err)
	}

	log.Info("configuration loaded")

	// Connect database
	db, err := database.New(cfg)
	if err != nil {
		log.Error("database connection failed")
		_ = log.Sync()
		return nil, fmt.Errorf("connect database: %w", err)
	}

	log.Info("database connected")

	// Initialize router
	r := router.New(cfg.AppEnv)

	// Configure HTTP server
	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &App{
		Config: cfg,
		Logger: log,
		DB:     db,
		Server: server,
	}, nil
}

// Run starts the HTTP server.
func (a *App) Run() error {
	a.Logger.Info("starting HTTP server")

	if err := a.Server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// Shutdown gracefully shuts down the application.
func (a *App) Shutdown(ctx context.Context) error {
	a.Logger.Info("shutting down application")

	if err := a.Server.Shutdown(ctx); err != nil {
		return err
	}

	if a.DB != nil {
		a.DB.Close()
	}

	if a.Logger != nil {
		_ = a.Logger.Sync()
	}

	return nil
}