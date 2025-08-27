package db

import (
	"fmt"
	"pastebin/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdb() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("pastedb.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = DB.AutoMigrate(&models.Pastes{})
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	fmt.Println("auto migrating successful")
	return nil
}
