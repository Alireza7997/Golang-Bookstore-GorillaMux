package models

import (
	"github.com/alireza/bookstore/internal/config"
	"github.com/jinzhu/gorm"
)

var database gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:""`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	database = *config.GetDB()
	database.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	database.NewRecord(book)
	database.Create(&book)

	return book
}

func GetAllBooks() []Book {
	var books []Book
	database.Find(&books)

	return books
}

func GetBookByid(id int64) (*Book, *gorm.DB) {
	var book Book
	database := database.Where("ID=?", id).Find(&book)

	return &book, database
}

func DeleteBook(id int64) Book {
	var book Book
	database.Where("ID=?", id).Delete(&book)

	return book
}
