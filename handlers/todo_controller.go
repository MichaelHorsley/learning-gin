package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTodoHandlers(r *gin.Engine) {
	r.GET("/ping", pingHandler)
	r.POST("/todo", createTodoHandler)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

func createTodoHandler(c *gin.Context) {
	var todo CreateTodoRequest
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"todoId": "1", "title": todo.Title})
}
