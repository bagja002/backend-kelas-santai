package main

import (
	"log"
	"project-kelas-santai/internal/config"
	"project-kelas-santai/internal/database"
	"project-kelas-santai/internal/models"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
	}

	// 2. Connect Database
	dsn := cfg.Database.DSN()
	database.ConnectDB(dsn)

	// 3. Drop Table
	err = database.DB.Migrator().DropTable(&models.Mentor{})
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}
	log.Println("Successfully dropped Mentor table")
}
