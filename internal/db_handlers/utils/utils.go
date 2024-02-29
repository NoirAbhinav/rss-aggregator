package db

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	models "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Migrate_models() (err error) {
	godotenv.Load(".env")
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(models.User{}, models.UserFeed{})
	if err != nil {
		log.Fatal(err)
		return errors.New("failed to migrate models")
	}
	fmt.Println("Migrated models")
	return
}

func GetDB() (db *gorm.DB, err error) {
	dsn := url.URL{
		User:   url.UserPassword(os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD")),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT")),
		Path:   os.Getenv("POSTGRES_DB_NAME"),
	}
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_USERNAME"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_DB_PORT"))
	db, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "rssagg.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return nil, errors.New("failed to connect database")
	}
	return
}
