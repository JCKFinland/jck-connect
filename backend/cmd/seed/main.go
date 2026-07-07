package main

import (
	"log"

	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
	"github.com/JCKFinland/jck-connect/backend/pkg/seed"
)

func main() {

	//----------------------------------------
	// Load configuration
	//----------------------------------------

	cfg := config.Load()

	//----------------------------------------
	// Connect database
	//----------------------------------------

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//----------------------------------------
	// Run seeds
	//----------------------------------------

	if err := seed.Run(db); err != nil {
		log.Fatal(err)
	}
}