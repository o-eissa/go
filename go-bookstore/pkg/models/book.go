package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `gorm:"" json:"author"`
	Publication string `gorm:"" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Find(&book).Delete(&book)
	return book
}

// func (b *Book) UpdateBook(id int) *Book {
// 	db.Save(&b)
// 	return b
// }
