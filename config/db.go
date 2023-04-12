package config

import (
	"errors"
	"log"
	"project-1-chapter-2/models"

	"gorm.io/gorm"
)

func GetAllBooks(book []models.Books) (allBooks []models.Books, err error) {
	db := GetDB()

	err = db.Find(&book).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book data not found")
		err = errors.New("error getting data")
		return
	}

	allBooks = append(allBooks, book...)
	return
}

func GetBookById(id int, book models.Books) (bookData models.Books, err error) {
	db := GetDB()

	err = db.First(&book, "id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	bookData = book
	return
}

func CreateBook(book models.Books) (newBook models.Books, err error) {
	var lastID int
	var lastBook models.Books
	db := GetDB()

	_ = db.Select("id").Last(&lastBook).Scan(&lastID)

	book = models.Books{
		ID:       uint(lastID) + 1,
		NameBook: book.NameBook,
		Author:   book.Author,
	}

	err = db.Create(&book).Error
	if err != nil {
		return
	}

	newBook = book
	return
}

func UpdateBook(id int, book models.Books) (updatedBook models.Books, err error) {
	db := GetDB()
	var findBook models.Books

	err = db.Where("id = ?", id).First(&findBook).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book not found")
		err = errors.New("book not found")
		return
	}

	err = db.Model(&book).Where("id = ?", id).Updates(models.Books{
		ID:        uint(id),
		NameBook:  book.NameBook,
		Author:    book.Author,
		CreatedAt: findBook.CreatedAt,
	}).Error

	if err != nil {
		log.Println("Error updating book data", err)
		err = errors.New("error updating book data")
		return
	}

	updatedBook = book
	return
}

func DeleteBook(id int) (err error) {
	db := GetDB()
	var book models.Books

	err = db.Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return
}
