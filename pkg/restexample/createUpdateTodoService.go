package restexample

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addTodo(context *gin.Context) {
	var newTodo todo

	//during binding if not all properties will be found to bind with todo struct then error will be thrown
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodoCompletion(context *gin.Context) {
	todo, err := findTodoByIdOrThrowError(context.Param("id"))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		todo.Completed = !todo.Completed
		context.Status(http.StatusNoContent)
	}
}
