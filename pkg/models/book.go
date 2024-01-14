package models

import (
	"github.com/Mahaveer86619/GO-Bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm: ""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {

	var getbook Book

	var db = db.Where("ID = ?", ID).Find(&getbook)

	return &getbook, db

}

func DeleteBook(ID int64) Book {

	var book Book

	db.Where("ID = ?", ID).Delete(&book)

	return book
}
