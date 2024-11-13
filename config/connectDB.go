package config

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data/referral_system.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	fmt.Println("Connected to database")
	if err := MigrateDB(); err != nil {
		log.Printf("Error during migration %v", err)
	}
	fmt.Println("Aplication started successfully")
}
