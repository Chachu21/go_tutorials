package main

import (
	"log"
	"os"

	"github.com/Chachu21/day-1/internal/database"
	"github.com/Chachu21/day-1/internal/handlers"
	"github.com/Chachu21/day-1/internal/models"
	"github.com/gin-gonic/gin"
	 "github.com/joho/godotenv"
)

func main (){
 err := godotenv.Load(".env")
 if err != nil{
  log.Fatalf("Error loading .env file: %s", err)
 }

	database.ConnectDB()
    database.DB.AutoMigrate(&models.Todo{})

    r := gin.Default()

// Define the root endpoint (GET /)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Todo API!",
		})
	})

	r.POST("/todos", handlers.CreateTodo)
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodoByID)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)	
}