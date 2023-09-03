package main

import (
	"database/sql"
	"go-structure-project/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	initDB()
	defer db.Close()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	// Setup routes
	routes.SetupRoutes(router, db)

	router.Run("localhost:8080")
}

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_books")
	if err != nil {
		panic(err.Error())
	}
}
