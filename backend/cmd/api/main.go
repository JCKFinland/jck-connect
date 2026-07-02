package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JCKFinland/jck-connect/backend/internal/app"
)

func main() {
	// Initialize application
	application, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	// Start HTTP server in a separate goroutine
	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("shutdown signal received")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Println("jck-connect stopped successfully")
}