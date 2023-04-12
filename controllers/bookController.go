package controllers

import (
	"fmt"
	"log"
	"net/http"
	"project-1-chapter-2/config"
	"project-1-chapter-2/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var books []models.Books

	allBooks, err := config.GetAllBooks(books)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Failed to get all datas %+v", err.Error()),
		})
		return
	}

	if allBooks == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": allBooks,
	})
}

func GetBookById(c *gin.Context) {
	bookID := c.Param("bookID")
	var bookData models.Books

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Invalid convert id")
		return
	}

	book, err := config.GetBookById(convBookID, bookData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convBookID),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var book models.Books

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook, err := config.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint("Failed to create book data because ", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	var book models.Books

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Failed to convert book id")
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook, err := config.UpdateBook(convBookID, book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Failed to update book data because %+v", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")

	convBookID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Failed to convert book id")
		return
	}

	err = config.DeleteBook(convBookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", convBookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d deleted successfully", convBookID),
	})
}
