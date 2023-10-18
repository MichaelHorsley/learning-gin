package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTodoHandlers(r *gin.Engine) {
	r.GET("/ping", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
