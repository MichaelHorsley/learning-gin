package main

import (
	"testing-gin-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handlers.RegisterTodoHandlers(r)

	r.Run("0.0.0.0:9050") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
