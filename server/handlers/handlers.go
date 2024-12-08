package handlers

import (
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

// handle database-related errors
func handleDatabaseError(c *gin.Context, err error, defaultMessage string) bool {
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": defaultMessage,
            "details": err.Error(),
        })
        return true
    }
    return false
}

// GetTodos - Retrieve all todos
func GetTodos(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, task, done FROM todos")
    if handleDatabaseError(c, err, "Failed to retrieve todos") {
        return
    }
    defer rows.Close()

    var todos []models.Todo
    for rows.Next() {
        var todo models.Todo
        if err := rows.Scan(&todo.ID, &todo.Task, &todo.Done); err != nil {
            handleDatabaseError(c, err, "Failed to scan todo")
            return
        }
        todos = append(todos, todo)
    }
    c.JSON(http.StatusOK, todos)
}

// CreateTodo - Add a new todo
func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := database.DB.Exec("INSERT INTO todos (task, done) VALUES (?, ?)", todo.Task, todo.Done)
    if handleDatabaseError(c, err, "Failed to create todo") {
        return
    }

    id, err := result.LastInsertId()
    if handleDatabaseError(c, err, "Failed to retrieve last insert ID") {
        return
    }

    todo.ID = int(id)
    c.JSON(http.StatusCreated, todo)
}

// ToggleTodo - Toggle the completion status of a todo
func ToggleTodo(c *gin.Context) {
    id := c.Param("id")
    _, err := database.DB.Exec("UPDATE todos SET done = NOT done WHERE id = ?", id)
    if handleDatabaseError(c, err, "Failed to toggle todo status") {
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
}

// DeleteTodo - Delete a todo by ID
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    _, err := database.DB.Exec("DELETE FROM todos WHERE id = ?", id)
    if handleDatabaseError(c, err, "Failed to delete todo") {
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}