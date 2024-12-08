package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id", handlers.ToggleTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	return router
}
