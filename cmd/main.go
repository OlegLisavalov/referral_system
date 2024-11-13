package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"referral_system/config"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	if err := config.MigrateDB(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Application started successfully")
}
