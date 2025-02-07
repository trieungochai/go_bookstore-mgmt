package models

import (
	"log"

	"github.com/trieungochai/go_bookstore-mgmt/pkg/config"
	"gorm.io/gorm"
)

// declares a global db variable that holds the database connection instance
var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init the database connection and auto-migrate the Book model
// This func is automatically called when the package is imported
func init() {
	config.ConnectDB()
	db = config.GetDB()
	// AutoMigrate func is used to automatically update the db schema
	// to match the Book struct.
	err := db.AutoMigrate(&Book{}).Error
	if err != nil {
		log.Fatalf("error migrating db: %v", err)
	}
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func GetBookById(Id int64) (*Book, error) {
	var book Book
	if err := db.Where("id = ?", Id).First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func Delete(Id int64) (*Book, error) {
	var book Book
	// First fetch the book to be deleted
	if err := db.Where("id = ?", Id).First(&book).Error; err != nil {
		return nil, err
	}

	// Then delete it
	if err := db.Where("id = ?", Id).Delete(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
