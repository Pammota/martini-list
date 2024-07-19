package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"martini-be/handlers"
	"martini-be/models"
)

func main() {
	var dbString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"))

	fmt.Println(dbString)

	db, err := gorm.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automigrate the tables
	db.AutoMigrate(&models.Item{})

	itemsHandler := handlers.CreateItemsHandler(db)

	router := gin.Default()

	// CORS Setup
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowOrigins = []string{"http://localhost:3000", "https://martini.gligor.dev"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "active",
		})
	})

	// Items CRUD
	router.GET("/items", itemsHandler.GetItems)
	router.GET("/items/:id", itemsHandler.GetItem)
	router.POST("/items", itemsHandler.CreateItem)
	router.PUT("/items/:id", itemsHandler.UpdateItem)
	router.DELETE("/items/:id", itemsHandler.DeleteItem)

	router.Run("0.0.0.0:8010")
}
