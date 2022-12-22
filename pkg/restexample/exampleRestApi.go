package restexample

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"Id"`
	Item      string `json:"Item"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{Id: "1", Item: "Clean Room", Completed: false},
	{Id: "2", Item: "Read Book", Completed: false},
	{Id: "3", Item: "Record Video", Completed: false},
}

func RunExampleRestApp() {
	var router = gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodoCompletion)
	router.POST("/todos", addTodo)
	router.Run("localhost:8080")
}
