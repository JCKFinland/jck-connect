package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/internal/container"
	"github.com/JCKFinland/jck-connect/backend/internal/router"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// App represents the application.
type App struct {
	config *config.Config
	db     *database.Database
	container *container.Container
	engine *gin.Engine
	server *http.Server
}

// New creates a new application using the default configuration.
func New() (*App, error) {
	return NewWithConfig(
		config.Load(),
	)
}

// NewWithConfig creates a new application using the provided configuration.
func NewWithConfig(
	cfg *config.Config,
) (*App, error) {

	//--------------------------------------------------
	// Connect database
	//--------------------------------------------------

	db, err := database.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	//--------------------------------------------------
	// Build dependency container
	//--------------------------------------------------

	c := container.New(
		cfg,
		db,
	)

	c.Compose()

	//--------------------------------------------------
	// Configure Gin
	//--------------------------------------------------

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	//--------------------------------------------------
	// Register routes
	//--------------------------------------------------

	router.Register(
		engine,
		c,
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
		container: c,
		engine: engine,
		server: server,
	}, nil
}

// Config returns the application configuration.
func (a *App) Config() *config.Config {
	return a.config
}

// Engine returns the HTTP handler.
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
	// Shutdown HTTP server
	//--------------------------------------------------

	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	//--------------------------------------------------
	// Close database
	//--------------------------------------------------

	if a.db != nil {
		a.db.Close()
	}

	return nil
}

// DB returns the application's database.
func (a *App) DB() *database.Database {
	return a.db
}

// Container returns the application's dependency container.
func (a *App) Container() *container.Container {
	return a.container
}

