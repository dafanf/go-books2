package routes

import (
	"database/sql"
	"go-structure-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Use anonymous functions to adapt controller functions to gin.HandlerFunc
	router.GET("/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})
	router.PATCH("/checkout", func(c *gin.Context) {
		controllers.CheckoutBook(c, db)
	})
	router.PATCH("/return", func(c *gin.Context) {
		controllers.ReturnBook(c, db)
	})
	router.POST("/books", func(c *gin.Context) {
		controllers.CreateBook(c, db)
	})
}
