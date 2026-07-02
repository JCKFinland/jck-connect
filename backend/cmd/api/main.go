package main

import (
	"log"

	"github.com/JCKFinland/jck-connect/backend/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal("failed to start application:", err)
	}

	_ = application

	log.Println("jck-connect backend started successfully")
}