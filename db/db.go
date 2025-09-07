package db

import (
	"fmt"
	"os"
	"pastebin/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdb() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong while loading the .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
