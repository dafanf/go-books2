package controllers

import (
	"go-structure-project/models"
	"go-structure-project/services"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context, db *sql.DB) {
	books, err := services.GetBooksFromDB(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func CheckoutBook(c *gin.Context, db *sql.DB) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}
	book, err := services.GetBookByID(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book out of stock"})
		return
	}
	book.Quantity--
	if err := services.UpdateBookQuantity(db, book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to update book quantity"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context, db *sql.DB) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}
	book, err := services.GetBookByID(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	book.Quantity++
	if err := services.UpdateBookQuantity(db, book); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to update book quantity"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context, db *sql.DB) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateBookInDB(db, newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}
