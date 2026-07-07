package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/internal/router"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// App represents the application.
type App struct {
	config *config.Config
	db     *database.Database
	engine *gin.Engine
	server *http.Server
}

// New creates a new application.
func New() (*App, error) {

	//--------------------------------------------------
	// Load configuration
	//--------------------------------------------------

	cfg := config.Load()

	//--------------------------------------------------
	// Connect database
	//--------------------------------------------------

	db, err := database.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	//--------------------------------------------------
	// Gin mode
	//--------------------------------------------------

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//--------------------------------------------------
	// Create Gin engine
	//--------------------------------------------------

	engine := gin.New()

	//--------------------------------------------------
	// Register routes
	//--------------------------------------------------

	router.Register(
		engine,
		cfg,
		db,
	)

	//--------------------------------------------------
	// HTTP server
	//--------------------------------------------------

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: engine,
	}

	return &App{
		config: cfg,
		db:     db,
		engine: engine,
		server: server,
	}, nil
}

// Config returns the application configuration.
func (a *App) Config() *config.Config {
	return a.config
}

// Engine returns the Gin engine.
func (a *App) Engine() *gin.Engine {
	return a.engine
}

// Run starts the HTTP server.
func (a *App) Run() error {

	log.Printf(
		"%s API starting on http://localhost:%s",
		a.config.AppName,
		a.config.AppPort,
	)

	err := a.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Shutdown gracefully shuts down the application.
func (a *App) Shutdown(
	ctx context.Context,
) error {

	//--------------------------------------------------
	// Stop HTTP server
	//--------------------------------------------------

	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	//--------------------------------------------------
	// Close database connection
	//--------------------------------------------------

	if a.db != nil {
		a.db.Close()
	}

	return nil
}