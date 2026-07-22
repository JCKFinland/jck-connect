package main

import (
	"log"

	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.Migrate(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	log.Println("Database migrations completed successfully.")
}
