package restexample

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	todo, err := findTodoByIdOrThrowError(context.Param("id"))

	if todo != nil {
		context.IndentedJSON(http.StatusOK, todo)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
}

func findTodoByIdOrThrowError(id string) (*todo, error) {
	for index, todo := range todos {
		if todo.Id == id {
			return &todos[index], nil
		}
	}
	return nil, errors.New("todo not found")
}
