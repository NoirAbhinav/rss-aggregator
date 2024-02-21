package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	models "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate_models() error {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return errors.New("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	return nil
}

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return nil, errors.New("failed to connect database")
	}
	return db, nil
}
