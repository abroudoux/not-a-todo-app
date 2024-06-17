package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo = []Todo{
	 {ID: 1, Title: "Go to the bank", Completed: false},
	{ID: 2, Title: "Buy some bread", Completed: true},
	{ID: 3, Title: "Learn Go", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodos)

	r.Run(":8080")
} 