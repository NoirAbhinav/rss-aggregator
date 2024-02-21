package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrate_models() error {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return errors.New("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return nil
}
