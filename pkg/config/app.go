package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Connect initializes the database connection.
func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
		return err
	}

	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("error connecting to database: %v", err)
		return err
	}

	log.Println("successfully connected to the database")
	return nil
}

// GetDB returns the active database connection.
func GetDB() *gorm.DB {
	return db
}
