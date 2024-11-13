package config

import (
	"fmt"
	"referral_system/models"
)

func MigrateDB() error {
	err := DB.AutoMigrate(&models.User{}, &models.RefCode{})
	if err != nil {
		fmt.Errorf("Failed to migrate database: %w\n", err)
		return err
	}
	fmt.Printf("Migrate database success\n")
	return nil
}
