package handlers

import (
	"net/http"

	"github.com/Chachu21/day-1/internal/database"
	"github.com/Chachu21/day-1/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    database.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func GetTodoByID(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
    var todo models.Todo
    id := c.Param("id")
    if err := database.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    database.DB.Delete(&models.Todo{}, id)
    c.Status(http.StatusNoContent)
}